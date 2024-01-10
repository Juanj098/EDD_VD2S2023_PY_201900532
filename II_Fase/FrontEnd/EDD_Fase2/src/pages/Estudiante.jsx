import { NavBarE } from "../components/navBarE";
import { useEffect, useState } from "react";

export function Estudiante() {
    const [Books, setBooks] = useState(null);
    const [Books1, setBooks1] = useState(null);
    const [Books2, setBooks2] = useState(null);

    const Peticion = async (c) => {
        try {
            const resp = await fetch("http://127.0.0.1:3000/Estudiante/Books", {
                method: 'POST',
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    "curso": parseInt(c)
                })
            });
    
            if (!resp.ok) {
                throw new Error(`HTTP error! Status: ${resp.status}`);
            }
    
            const libros = await resp.json();
            const libro = libros.LibrosAceptados
            setBooks(libro);
        } catch (error) {
            console.error("Error fetching data:", error);
        }
    }
    const Peticion1 = async (c) => {
        try {
            const resp = await fetch("http://127.0.0.1:3000/Estudiante/Books", {
                method: 'POST',
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    "curso": parseInt(c)
                })
            });
    
            if (!resp.ok) {
                throw new Error(`HTTP error! Status: ${resp.status}`);
            }
    
            const libros = await resp.json();
            const libro = libros.LibrosAceptados
            setBooks1(libro);
        } catch (error) {
            console.error("Error fetching data:", error);
        }
    }  
    const Peticion2 = async (c) => {
        try {
            const resp = await fetch("http://127.0.0.1:3000/Estudiante/Books", {
                method: 'POST',
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    "curso": parseInt(c)
                })
            });
    
            if (!resp.ok) {
                throw new Error(`HTTP error! Status: ${resp.status}`);
            }
    
            const libros = await resp.json();
            const libro = libros.LibrosAceptados
            setBooks2(libro);
        } catch (error) {
            console.error("Error fetching data:", error);
        }
    }
        

  return (
    <>
      <NavBarE />
      <div className="container">
        <h1>Portal Estudiante</h1>
        <div className="containerCards">
            <div className="card">
                <h1>{localStorage.getItem('c1')}</h1>
                <div>
                    <a onClick={() => Peticion(localStorage.getItem('c1'))}>Libros</a>
                    <a>Publicaciones</a>
                </div>
                <div>
                    {Books?.map((libro, index) => (
                        <ul key={index}>
                            <li>{libro.Name}</li>
                        </ul>
                    ))}
                </div>
            </div>
            <div className="card">
                <h1>{localStorage.getItem('c2')}</h1>
                <div>
                    <a onClick={() => Peticion(localStorage.getItem('c2'))}>Libros</a>
                    <a>Publicaciones</a>
                </div>
                <div>
                    {Books1?.map((libro, index) => (
                        <ul key={index}>
                            <li>{libro.Name}</li>
                        </ul>
                    ))}
                </div>
            </div>
            <div className="card">
                <h1>{localStorage.getItem('c3')}</h1>
                <div>
                    <a onClick={() => Peticion(localStorage.getItem('c3'))}>Libros</a>
                    <a>Publicaciones</a>
                </div>
                <div>
                    {Books2?.map((libro, index) => (
                        <ul key={index}>
                            <li>{libro.Name}</li>
                        </ul>
                    ))}
                </div>
            </div>
        </div>
      </div>
    </>
  );
}
