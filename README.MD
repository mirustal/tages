# File Service

Этот сервис реализует функциональность для загрузки, скачивания и получения списка файлов. Используется gRPC для реализации следующих методов: `UploadFile`, `ListFiles`, и `DownloadFile`.

Для запуска в папке, где лежит docker-compose файл выполнить команду
** docker-compose up --build

## Сервис: `Requester`
Для удобства проверки работоспособности был написан сервис requester

Программа `requester` позволяет пользователю взаимодействовать с сервисом FileService и выполнять три основных операции: скачивание файла, получение списка файлов, и загрузка файла на сервер. Программа работает с gRPC сервисом и взаимодействует с файловой системой для сохранения и чтения файлов.

### Основные функции:
1. **Скачать файл**: Клиент запрашивает файл по имени и сохраняет его в локальной папке `download`.
2. **Получить список файлов**: Клиент запрашивает список всех файлов, доступных на сервере, и выводит их  (имя, дата создания, дата обновления).
3. **Загрузить файл**: Клиент загружает файл с локального диска на сервер 

### Пример запуска программы

Для запуска программы `requester`, выполните следующие шаги:

1. Откройте терминал и выполните команду для запуска программы:

    ```bash
    go run file-service/cmd/requester/main.go
    ```

2. После запуска программы вам будет предложено выбрать одно из действий:
    ```
    Выберите действие:
    1 - Скачать файл
    2 - Получить список файлов
    3 - Загрузить файл
    ```

3. Выберите одно из действий, введя соответствующий номер (1, 2 или 3) и нажав Enter.


## Сервис: `FileService`

### Методы

#### 1. UploadFile (Загрузка файла)

**RPC метод:** `UploadFile(stream UploadRequest) returns (UploadResponse)`

Используется для загрузки файлов на сервер по частям.

**Запрос: `UploadRequest`**
- `bytes file_chunk` - Часть файла в бинарном формате.
- `string file_name` - Имя файла, который загружается.

**Ответ: `UploadResponse`**
- `string file_name` - Имя загруженного файла.
- `string message` - Сообщение об успешной загрузке.

**Пример использования:**
```
grpcurl -plaintext -import-path /Users/mirustal/Documents/project/go/tages/file-service/api/file -proto file.proto -d '{

  "file_name": "example3.jpg",

  "file_chunk": "'$(base64 -i /Users/mirustal/Desktop/example2.jpg)'"

}' localhost:9002 file.FileService/UploadFile
```

---

#### 2. ListFiles (Список файлов)

**RPC метод:** `ListFiles(ListFilesRequest) returns (ListFilesResponse)`

Возвращает список файлов, доступных на сервере.

**Запрос: `ListFilesRequest`**
- Нет параметров.

**Ответ: `ListFilesResponse`**
- `repeated FileMetadata files` - Список метаданных файлов.

**Структура `FileMetadata`:**
- `string file_name` - Имя файла.
- `string created_at` - Дата создания файла (в формате строки).
- `string updated_at` - Дата последнего обновления файла (в формате строки).

**Пример использования:**
```
grpcurl -plaintext -import-path /Users/mirustal/Documents/project/go/tages/file-service/api/file -proto file.proto -d '{}' localhost:9002 file.FileService/ListFiles
```

---

#### 3. DownloadFile (Скачивание файла)

**RPC метод:** `DownloadFile(DownloadRequest) returns (DownloadResponse)`

Используется для скачивания файла с сервера.

**Запрос: `DownloadRequest`**
- `string file_name` - Имя файла, который нужно скачать.

**Ответ: `DownloadResponse`**
- `bytes file_chunk` - Часть файла в бинарном формате.
- `string file_name` - Имя файла.

**Пример использования:**
```
grpcurl -plaintext -import-path /Users/mirustal/Documents/project/go/tages/file-service/api/file -proto file.proto -d '{ 

  "file_name": "example.jpg"

}' -format json localhost:9002 file.FileService/DownloadFile
```

---

