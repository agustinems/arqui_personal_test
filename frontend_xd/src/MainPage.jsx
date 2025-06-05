import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './MainPage.css';

function MainPage() {
  const [activities, setActivities] = useState([]);
  const [error, setError] = useState(null);
44qji1-codex/crear-frontend-página-main
  const [selected, setSelected] = useState(null);

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

44qji1-codex/crear-frontend-página-main
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
44qji1-codex/crear-frontend-página-main
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
          </div>
        ))}
      </div>
 main
    </div>
  );
}

export default MainPage;
