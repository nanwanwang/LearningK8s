FROM mcr.microsoft.com/dotnet/core/sdk:2.1 AS build

WORKDIR /app
RUN dotnet restore

# COPY docker-console/docker-console.csproj ./docker-console/
# WORKDIR /app/docker-console
# RUN dotnet restore

WORKDIR /app
# COPY docker-console/. ./docker-console/
# WORKDIR /app/docker-console
RUN dotnet publish -c Release -o out


FROM mcr.microsoft.com/dotnet/core/runtime:2.1 AS runtime
WORKDIR /app
COPY --from=build /app/out ./
ENTRYPOINT ["dotnet","docker-console.dll"]



