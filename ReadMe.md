# BACKEND DEVELOPMENT IN GO
Backend development in Go is very easy when you understande the basics, the readme file is just a basic explanation of my understanding. i will definitely get better with time

## GETTING STARTED WITH THE CODESPACE
 - This is just the basic getting started with go env. 
 1. create a main file - and test if the file is working by initializing, building, and running. 
 2. create the go.mod file using the cmd prompt below
    *  make sure you are in the right directory(folder) where your file is
        you can cd in that directory using this command.
    ```bash pws
    cd ./(directoryname)
    ```
    * you need to build an initialization file for go to use
        you can do this by typing the code below
        ```bash pws
        go mod init
        ```
    * Now we can build our program.
        use the code below.
        ```bash pws
        go build
        ```
        afte the program is built you will find the name of the program with an ```exe``` extention that is the program you are going to run
    * Lets run our program
        use the code below
        ```bash pws
        ./(filename).exe

## GETING THE FRAMEWORKS
- This is just a rough explanation of my understanding from the beginning to the current place so far.
to have a configuration like this, you will need to have certain modules/framework imported into you working directory.
this frameworks can be dowloaded from git hub and integrated as follows
1. ```bash pws
        go get github.com/<frameworkwriter>/<framework>
        once the conmmand above is executed, the frame work in then dowlaonded into the working directory.  
    ```
2. ```bash psw
        after you have downloaded the frame work you will have to run the execution command below to create a vendor
        go mod vender
        go mod tidy
        now you have to check if the mod file in the working directory has the downloaded frameworks in it. it should be located at the required section of the go.mod file
        if not arranged well, run the ```go mod vendor``` after the ```go mod tidy```
    ```
3. repeat step one and two for every framework you want to integrate

## BUILDING OUR FIRST ROUTE
    to build our first route without security features, we will need some three frameworks, without the inbuilt frameworks that come with go.
    this frameworks are external and we will have to complete the getting the frameworks section
    listed below is are the initial frameworks to get
        1. dotenv reader which is from "joho/gedotenv" joho is the writer name
        2. chi router which from go-chi/chi - here you will need the updated version of the frame work, hence you should add the version - go-chi/chi/v?
        3. cors from the same go-chi/cors - we can't build a framework without adding a cross origin