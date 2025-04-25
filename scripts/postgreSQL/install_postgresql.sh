#!/bin/bash

# Installer PostgreSQL
# Version: 1.2
# Author: Eugene Zhelak

# Что делает скрипт?

# Этот скрипт автоматизирует установку и настройку PostgreSQL на Linux (например, Raspberry Pi). Он выполняет:

# Обновление системы
# Удаление старых версий PostgreSQL (если есть)
# Установку PostgreSQL и дополнительных компонентов
# Настройку пароля для пользователя postgres
# Создание административного пользователя БД (имеет все права кроме суперпользователя шаг 5)
# Настройку удаленного доступа
# Создание основной базы данных

# Как пользоваться:

# Сохраните скрипт как файл (например, install_postgres.sh)
# Дайте права на выполнение:

# chmod +x install_postgres.sh
# Запустите с правами root (можно с параметрами):

# sudo ./install_postgres.sh
# Дополнительные параметры (необязательно):

# -pp или --postgres-password - пароль для пользователя postgres (по умолчанию: postgres)
# -nd или --db-name - имя основной БД (по умолчанию: app_db)
# -ua или --admin-user - имя администратора БД (по умолчанию: admin)
# -pa или --admin-password - пароль администратора (по умолчанию: password)
# Пример с параметрами:

# sudo ./install_postgres.sh --postgres-password "postgres" --db-name "app_db" --admin-user "admin" --admin-password "password"
# =================================================================================================

set -e  # Прерывать выполнение при ошибках

# Конфигурационные переменные по умолчанию
POSTGRES_PASSWORD="postgres" # postgres (root)
DB_NAME="app_db"
ADMIN_USER="admin"
ADMIN_PASSWORD="password"

#Парсинг аргументов
while [[ $# -gt 0 ]]; do
    case "$1" in
        -pp|--postgres-password)
            [[ -n "$2" ]] && POSTGRES_PASSWORD="$2"
            shift 2
            ;;
        -nd|--db-name)
            [[ -n "$2" ]] && DB_NAME="$2"
            shift 2
            ;;
        -ua|--admin-user)
            [[ -n "$2" ]] && ADMIN_USER="$2"
            shift 2
            ;;
        -pa|--admin-password)
            [[ -n "$2" ]] && ADMIN_PASSWORD="$2"
            shift 2
            ;;
        *)
            # Игнорируем все остальные аргументы
            shift
            ;;
    esac
done

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Функции для удобства
info() { echo -e "${YELLOW}[INFO] $1${NC}"; }
success() { echo -e "${GREEN}$1${NC}"; }
step() { echo -e "${GREEN}[$1/7] $2${NC}"; }
error() { echo -e "${RED}[ERROR] $1${NC}"; exit 1; }

# Проверка прав root
if [[ $EUID -ne 0 ]]; then
    error "Ошибка: Этот скрипт требует root прав используйте sudo!"
fi

info "Начало установки PostgreSQL..."

# 1. Обновление системы
step "1" "Обновление пакетов..."
sudo apt update
sudo apt upgrade -y

# 2. Удаление существующего PostgreSQL (если есть)
step "2" "Очистка предыдущих установок PostgreSQL..."

# Проверка установлен ли PostgreSQL
if dpkg -l | grep -q postgresql; then
    info "Обнаружена установленная версия PostgreSQL, начинаю удаление..."
    sudo systemctl stop postgresql || true
    sudo apt-get remove --purge -y postgresql*
    sudo apt-get autoremove -y
    sudo apt-get autoclean

    # Дополнительная очистка оставшихся файлов
    sudo rm -rf /var/lib/postgresql/
    sudo rm -rf /var/log/postgresql/
    sudo rm -rf /etc/postgresql/
    sudo rm -rf /etc/postgresql-common/
    sudo rm -rf /var/cache/postgresql/
    info "Старые версии PostgreSQL полностью удалены"
else
    info "PostgreSQL не установлен, пропускаю шаг удаления"
fi

# 3. Установка PostgreSQL
step "3" "Установка PostgreSQL..."
sudo apt install -y postgresql postgresql-contrib

# Проверка версии
PG_VERSION=$(psql --version | awk '{print $3}' | cut -d. -f1)
info "Установлена PostgreSQL версии ${PG_VERSION}"

# 4. Настройка пользователя postgres
step "4" "Настройка пользователя postgres..."
sudo -u postgres psql -c "ALTER USER postgres WITH PASSWORD '${POSTGRES_PASSWORD}';"
info "Пароль для postgres установлен: '${POSTGRES_PASSWORD}'"

# 5. Создание административного пользователя назначение прав
step "5" "Создание пользователя ${ADMIN_USER}..."
sudo -u postgres psql <<-EOSQL
    CREATE USER ${ADMIN_USER} WITH 
        NOSUPERUSER
        CREATEDB
        CREATEROLE
        INHERIT
        LOGIN
        REPLICATION
        BYPASSRLS
        PASSWORD '${ADMIN_PASSWORD}';
    GRANT ALL PRIVILEGES ON DATABASE postgres TO ${ADMIN_USER};
EOSQL
info "Создан пользователь ${ADMIN_USER} с паролем '${ADMIN_PASSWORD}'"

# 6. Настройка удаленного доступа
step "6" "Настройка сетевого доступа..."
CONF_FILE="/etc/postgresql/${PG_VERSION}/main/postgresql.conf"
HBA_FILE="/etc/postgresql/${PG_VERSION}/main/pg_hba.conf"

# Разрешить подключения со всех адресов
sudo sed -i "s/#listen_addresses = 'localhost'/listen_addresses = '*'        /" "$CONF_FILE"

# Добавить правило доступа
echo "host    all             all             0.0.0.0/0             md5" | sudo tee -a "$HBA_FILE" > /dev/null

# 7. Перезапуск и проверка PostgreSQL
step "7" "Перезапуск и настройка службы..."
sudo systemctl restart postgresql
sudo systemctl enable postgresql >/dev/null 2>&1

# Проверка статуса службы
if ! systemctl is-active --quiet postgresql; then
    error "PostgreSQL не запустился после перезагрузки"
else
    info "PostgreSQL успешно запущен и добавлен в автозагрузку"
fi

# Создание основной БД
info "Создание основной базы данных ${DB_NAME}..."
sudo -u postgres createdb "${DB_NAME}" --owner="${ADMIN_USER}"
info "Создана база данных: ${DB_NAME} (владелец: ${ADMIN_USER})"

# Итоговое сообщение
echo -e "\n${GREEN}[Готово] PostgreSQL успешно установлен и настроен!${NC}"
echo -e "Основные данные для подключения:"
echo -e "  Хост: localhost или IP вашего сервера"
echo -e "  Пользователь: postgres (root)"
echo -e "  Пароль: ${POSTGRES_PASSWORD}"
echo -e "  Порт: 5432"
echo -e "\nДополнительные данные:"
echo -e "  Основная БД: ${DB_NAME}"
echo -e "  Административный пользователь: ${ADMIN_USER}"
echo -e "  Пароль администратора: ${ADMIN_PASSWORD}"