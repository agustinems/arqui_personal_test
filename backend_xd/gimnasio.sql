DROP DATABASE IF EXISTS gimnasio;
CREATE DATABASE gimnasio;
USE gimnasio;

-- (pegá después el script completo que te di con tablas + inserts)
-- Crear base y usarla
CREATE DATABASE IF NOT EXISTS gimnasio;
USE gimnasio;

-- Tabla de usuarios
CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    es_admin BOOLEAN NOT NULL DEFAULT FALSE
);

-- Tabla de categorías
CREATE TABLE categorias (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL
);

-- Tabla de actividades
CREATE TABLE actividades (
    id INT AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(100) NOT NULL,
    descripcion TEXT,
    foto_url TEXT,
    categoria_id INT,
    cupo INT NOT NULL,
    duracion INT,
    instructor_nombre VARCHAR(100),
    instructor_email VARCHAR(100),
    instructor_descripcion TEXT,
    FOREIGN KEY (categoria_id) REFERENCES categorias(id)
);

-- Tabla de clases
CREATE TABLE clases (
    id INT AUTO_INCREMENT PRIMARY KEY,
    actividad_id INT NOT NULL,
    dia DATE NOT NULL,
    horario TIME NOT NULL,
    FOREIGN KEY (actividad_id) REFERENCES actividades(id)
);

-- Tabla de inscripciones
CREATE TABLE inscripciones (
    id INT AUTO_INCREMENT PRIMARY KEY,
    usuario_id INT NOT NULL,
    clase_id INT NOT NULL,
    fecha_inscripcion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
    FOREIGN KEY (clase_id) REFERENCES clases(id)
);

-- Usuarios de prueba
INSERT INTO usuarios (nombre, email, password_hash, es_admin) VALUES
('Admin Test', 'admin@test.com', '$2a$12$axgziaZApgoE4cyUXEqZm.KvF7pgr4CTRuAIWswcy9aZRhEKXg/w2', TRUE), -- contraseña: 123456
('Juan Pérez', 'juan@socio.com', '$2a$12$axgziaZApgoE4cyUXEqZm.KvF7pgr4CTRuAIWswcy9aZRhEKXg/w2', FALSE); -- contraseña: 123456

-- Categorías
INSERT INTO categorias (nombre) VALUES 
('Funcional'), 
('Spinning'), 
('MMA');

-- Actividades
INSERT INTO actividades (titulo, descripcion, foto_url, categoria_id, cupo, duracion, instructor_nombre, instructor_email, instructor_descripcion) VALUES
('Clase Funcional', 'Entrenamiento completo de alta intensidad.', 'https://via.placeholder.com/150', 1, 20, 60, 'Ana Torres', 'ana@fit.com', 'Instructora certificada en fitness funcional.'),
('Spinning Avanzado', 'Clases de ciclismo indoor para quemar calorías.', 'https://via.placeholder.com/150', 2, 15, 45, 'Carlos Spin', 'carlos@bike.com', 'Instructor con 5 años de experiencia.'),
('Entrenamiento MMA', 'Técnicas básicas de artes marciales mixtas.', 'https://via.placeholder.com/150', 3, 10, 90, 'Luis Lucha', 'luis@mma.com', 'Experto en jiu-jitsu y muay thai.');

-- Clases de esas actividades
INSERT INTO clases (actividad_id, dia, horario) VALUES
(1, '2025-06-01', '10:00:00'),
(1, '2025-06-03', '10:00:00'),
(2, '2025-06-02', '18:00:00'),
(3, '2025-06-05', '19:30:00');

-- Inscripción de prueba (socio Juan se inscribió a la clase funcional del 1/6)
INSERT INTO inscripciones (usuario_id, clase_id) VALUES
(2, 1);

