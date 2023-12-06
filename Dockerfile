    # Utiliza la imagen oficial de MongoDB como base
FROM mongo:7.0.4-jammy

# Copia un script personalizado para inicializar la base de datos
COPY db/init-database.js /docker-entrypoint-initdb.d/


# Define variables de entorno para la configuración
ENV MONGO_INITDB_ROOT_USERNAME=admin
ENV MONGO_INITDB_ROOT_PASSWORD=12345
ENV MONGO_INITDB_DATABASE=events

# Exponer el puerto de MongoDB
EXPOSE 27017

# Comando por defecto para iniciar MongoDB
CMD ["mongod"]



