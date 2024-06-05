# Використовуємо образ Node.js на базі Alpine Linux
FROM node:20.14.0-alpine

# Встановлюємо робочу директорію в /app
WORKDIR /app

# Копіюємо файли package.json та package-lock.json для встановлення залежностей
COPY package*.json ./

# Встановлюємо залежності за допомогою npm
RUN npm install

# Копіюємо решту файлів з вашого проекту в робочу директорію образу
COPY . .

# Виконуємо команду для запуску вашого застосунку
CMD ["node", "app.js"]

