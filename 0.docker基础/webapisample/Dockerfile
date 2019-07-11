FROM mcr.microsoft.com/dotnet/core/sdk:3.0.100-preview6 AS build
WORKDIR /app

COPY ./*.csproj  ./
WORKDIR /app
RUN dotnet restore 

WORKDIR /app
COPY  /. ./
WORKDIR /app
RUN dotnet publish -c Release -o ./out


FROM mcr.microsoft.com/dotnet/core/aspnet:3.0.0-preview6 AS runtime
WORKDIR /app
COPY --from=build /app/out .
ENTRYPOINT [ "dotnet","webapisample.dll" ]