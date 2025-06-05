import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './MainPage.css';

function MainPage() {
  const [activities, setActivities] = useState([]);
  const [error, setError] = useState(null);
  const [editing, setEditing] = useState(null);
  const [formData, setFormData] = useState({
    titulo: '',
    descripcion: '',
    dia: '',
    horario: '',
    cupo: 0,
    imagen: '',
  });
  const fetchAll = () => {
  };

  useEffect(() => {
    fetchAll();
  const startEdit = (act) => {
    setEditing(act.id);
    setFormData({
      titulo: act.titulo,
      descripcion: act.descripcion,
      dia: act.dia,
      horario: act.horario,
      cupo: act.cupo,
      imagen: act.imagen,
    });
  };

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const saveEdit = () => {
    axios.put(`http://localhost:8000/actividad/${editing}`, formData)
      .then(() => {
        setEditing(null);
        fetchAll();
      })
      .catch(() => {
        setError('No se pudo actualizar la actividad');
      });
  };

            <button type="button" onClick={() => startEdit(act)}>
              Editar
            </button>
      {editing && (
        <div className="detail">
          <h2>Editar actividad</h2>
          <input name="titulo" value={formData.titulo} onChange={handleChange} />
          <input name="descripcion" value={formData.descripcion} onChange={handleChange} />
          <input name="dia" value={formData.dia} onChange={handleChange} />
          <input name="horario" value={formData.horario} onChange={handleChange} />
          <input name="imagen" value={formData.imagen} onChange={handleChange} />
          <input name="cupo" type="number" value={formData.cupo} onChange={handleChange} />
          <button type="button" onClick={saveEdit}>Guardar</button>
          <button type="button" onClick={() => setEditing(null)}>Cancelar</button>
        </div>
      )}
  const [selected, setSelected] = useState(null);


44qji1-codex/crear-frontend-página-main
  const [selected, setSelected] = useState(null);

 main

 main
  useEffect(() => {
    axios.get('http://localhost:8000/actividades')
      .then((response) => {
        setActivities(response.data);
        setError(null);
      })
      .catch(() => {
        setError('No se pudieron cargar las actividades');
      });
  }, []);

 9evkb5-codex/crear-frontend-página-main

44qji1-codex/crear-frontend-página-main
main
  const fetchActivity = (id) => {
    axios.get(`http://localhost:8000/actividad/${id}`)
      .then((response) => {
        setSelected(response.data);
        setError(null);
      })
      .catch(() => {
        setError('No se pudo obtener la actividad');
      });
  };

 9evkb5-codex/crear-frontend-página-main


 main
main
  return (
    <div className="main-container">
      <h1>Actividades</h1>
      {error && <p className="error">{error}</p>}
      <div className="cards">
        {activities.map((act) => (
          <div key={act.id} className="card">
            <img src={act.imagen} alt={act.titulo} />
            <h2>{act.titulo}</h2>
            <p>{act.descripcion}</p>
            <p>
              <strong>Dia:</strong> {act.dia} - <strong>Horario:</strong> {act.horario}
            </p>
            <p>
              <strong>Cupo:</strong> {act.cupo}
            </p>
 9evkb5-codex/crear-frontend-página-main

44qji1-codex/crear-frontend-página-main
 main
            <button type="button" onClick={() => fetchActivity(act.id)}>
              Ver detalles
            </button>
          </div>
        ))}
      </div>
      {selected && (
        <div className="detail">
          <h2>Detalle de actividad</h2>
          <p>
            <strong>Título:</strong> {selected.titulo}
          </p>
          <p>{selected.descripcion}</p>
          <p>
            <strong>Día:</strong> {selected.dia} - <strong>Horario:</strong>{' '}
            {selected.horario}
          </p>
          <p>
            <strong>Cupo:</strong> {selected.cupo}
          </p>
        </div>
      )}
 9evkb5-codex/crear-frontend-página-main

          </div>
        ))}
      </div>
 main
 main
    </div>
  );
}

export default MainPage;
