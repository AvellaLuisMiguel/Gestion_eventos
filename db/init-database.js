// Especifica la base de datos directamente
db = db.getSiblingDB('events');

// Crea un usuario para la base de datos 'events'
db.createUser({
  user: 'admin',
  pwd: '12345',
  roles: [{ role: 'readWrite', db: 'events'}]
});


// Insertar documentos en la colección 'events'
db.events.insertMany([
  {
    id_event: 1,
    name_event: 'Evento 1',
    type_event: 'Tipo A',
    description_event: 'Descripción del Evento 1',
    date_event: '2023-01-01',
    state_event: 1,
  },
  {
    id_event: 2,
    name_event: 'Evento 2',
    type_event: 'Tipo B',
    description_event: 'Descripción del Evento 2',
    date_event: '2023-01-01',
    state_event: 2,
  },
]);

print('Inicialización completada------------------------------------------------------------------------------------------------------');