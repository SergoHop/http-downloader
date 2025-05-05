# HTTP Загрузчик файлов

Простая утилита для загрузки файлов по HTTP, написанная на Go без использования внешних фреймворков.

## Возможности

- Базовая загрузка файлов
- Отображение прогресса загрузки
- Возобновление прерванных загрузок
- Параллельная загрузка (несколько соединений)

- # Простая загрузка
./http-downloader http://example.com/file.zip

# Загрузка с возможностью докачки
./http-downloader http://example.com/file.zip output.zip resume

# Параллельная загрузка (4 потока)
./http-downloader http://example.com/file.zip output.zip parallel 4
