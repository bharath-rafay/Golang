# Project
An Golang application to process the records and write them in the format of YAML or JSON based on the Environment varible ("FORMAT") onto a File

>**INPUT FILES**

inp.txt
Rohan,12,M,[cricket,football],5.9,50
Rohit,11,M,[football],5.8,51
Keerthi,13,F,[badminton,table tennis],5.5,45
Rohini,12,M,[tennis],5.4,44
Rakesh,12,M,[cricket],5.9,55
Vinay,12,M,[chess,carrom],5.7,52
Neha,12,F,[volleyball],5.7,40

>**Key Components of the Code **

**Getinfo(slice,file)** - sport.go 
      reads the text from file and returns a slice of struct with all the details
**DetailWriter interface** -sport.go 
      Its an interface used to create either yaml or json file of the input file
**CreateFile** -sport.go 
      this function is used to create the yaml or json file by passing the interface structs
      


**logger**
    logger is another folder where we define all logging functions 


**Output**
A file with all the records processed in the fomat of .yaml or .json



**External Packages used **

  *)Viper - To access Environment variables 
  

