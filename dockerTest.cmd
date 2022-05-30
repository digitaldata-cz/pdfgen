@ECHO OFF
IF [%1]==[build] CALL :BUILD
IF [%1]==[run] CALL :RUN
IF [%1]==[all] CALL :ALL
EXIT /B 1

:ALL
 CALL :BUILD
 CALL :RUN
EXIT /B 0

:BUILD
 CALL docker build -t pdfgen .
 ECHO "build done"
EXIT /B 0

:RUN
 CALL docker run -it -p 50051:50051 -e PG_IP=0.0.0.0 -e PG_PORT=50051 pdfgen
EXIT /B 0


