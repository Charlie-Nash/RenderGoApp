-- Crear la tabla de vendedores
CREATE TABLE vendedores (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL
);

-- Insertar algunos datos de prueba
INSERT INTO vendedores (nombre) VALUES
('Juan Pérez'),
('María García'),
('Carlos López'),
('Ana Sánchez'),
('Luis Rodríguez');
