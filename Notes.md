# 自己的Docker筆記

## **Docker image**

---
說明：  
類似映像檔，但是更輕量化


## **Docker container**

---
說明：  
映像檔建立後的實體(容器)，稱回container



## **Dockerfile**

---
說明：  
以一個服務為基礎，再透過文件裡面的指令建立相關功能。最終透過 docker build 指令，可以把Dockerfile打包成 Docker Image。簡單來說變成屬於自己專用的 Image。<br>


### *Dockerfile 文件內參數：*

* FROM => 基本映像檔，必需(ex: ubuntu:22.4, nginx:1.14, php:latest)
* LABEL => 增加一些原始數據(資訊)，非必須(ex: 作者, 版號, 基礎描述)
* MAINTAINER => 維護者名稱，非必須
* COPY => 複製目前資料夾的檔案 至 映像檔內，非必須(ex: copy /<本地資料夾> /<新環境資料夾>)
* ADD => 複製目前資料夾的檔案 至 映像檔內，聽說可以自動解壓縮和url下載，沒試過，非必須
* WORKDIR => 切換工作目錄，非必須(ex: /home)
* RUN => 預設要執行的指令，會常用到， 類似ubuntu基本image並沒有vim，就可以先執行安裝vim套件。
       簡單來說，container內的下指令的動作也寫入Dockerfile，
       非必須(ex: RUN apt-get install vim && apt-get install <某某套件名稱>)
* EXPOSE => 是在告訴Docker，某服務是啟動在某個port。類似聲明而已，並不會真的啟動，非必須(ex: EXPOSE 80/tcp)

* CMD => CMD 為command縮寫，對應 docker run [OPTIONS] IMAGE [COMMAND] [ARG...] 中的 [COMMAND]，一個Dockerfile只能一個CMD
       必須(ex: ["/bin/bash"])



### *Docker build 指令*

```
指令格式： docker build [--option] <Image名稱> <Dockerfile> [--option]  
EX: docker build -t <image-name> <Dockerfile路徑> --no-cache
```



## **Docker Compose yml 文件**

---

說明：  
簡單來如，太多服務要用就用yml檔，一次就呼叫所有image或是Dockerfile，然後把所有服務的參數和服務間的連結設定好  
比如: 伺服器+腳本語言+資料庫+其他哩哩摳摳的工具
<br>

### *Docker Compose yml 文件內參數：*

* version => yaml格式版本，記錄當下是3.9版本，必須(ex: "3.9")
* services => 對應container，一個services區塊就是一個容器的建立參數，所以可以多個services，每個services的命名就是container-name，必須(ex: app: , database: , web: ,)
* volumes => 外部的資料夾和容器的資料夾 兩位置連結 (ps 這樣備份或存入某東西感覺很方便)
* ports => 對外PORT號(ex: "8081:81")
* environment => 環境變數，這個需參照各image的網路上文件說明(ex: 如Database的root帳密之類的)
* networks => 可以設定一個網路給某個容器，但是容器建立時候通常也會建立一個簡單的網路使用

### *Docker Compose 指令*

```
指令格式： docker-compose [--option] [--option]

EX: 
docker-compose -f docker-compose.yml up -d  
建立容器並且在背景執行容器

docker-compose -f stop
停止執行容器

docker-compose -f docker-compose.yml start   
重新執行容器(在背景運作)

docker-compose -f down
停止並刪除容器

```


## **Docker run 指令**

---
說明：  
用來執行image，建立docker container


重要 OPTIONS：  
* -i & -t 搭配使用建立一個shell給user
* -p 虛擬環境:本地環境 兩個port號
* -v 虛擬環境:本地環境 兩者的資料夾互通
* --name 幫容器取名

[COMMAND] => 通常是寫 bash || /bin/bash 不然沒辦法使用 shell

``` 
指令格式： docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

EX: docker run -i -t -p 80:80 --name="container-name" -v /data:/data <image-name> bash
```



## **Docker exec 指令**

---
說明：  
用來進入已建立的docker container

``` 
指令格式： docker exec [OPTIONS] CONTAINER COMMAND [ARG...]

EX: docker exec -it <container-name> bash
```
