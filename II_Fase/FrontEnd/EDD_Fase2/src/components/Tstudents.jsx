import { NavBar } from "./navbar"
import { useState , useEffect} from "react"
export function TablaStudents (){

    const [students, Setstudents] = useState(null)
    useEffect(()=>{
        fetch('http://127.0.0.1:3000/Admin/GetStudents')
        .then((response)=>response.json())
        .then((data) => Setstudents(data))
    },[])

    console.log(students)
    
    return(
        <>
        <NavBar/>
        <div className="container">
            <table className="Testudiante">
                <thead>
                    <th key="llave-header" >Llave</th>
                    <th key="carnet-header">Carnet</th>
                    <th key="nombre-header">Nombre</th>
                    <th key="password-header">Password</th>
                    <th key="cursos-header">Cursos</th>
                </thead>
                <tbody>
                {students?.map((estudiante)=>(
                    <tr key={estudiante.Llave}>
                        <td>{estudiante.Llave}</td>
                        <td>{estudiante.Dato.Carnet}</td>
                        <td>{estudiante.Dato.Name}</td>
                        <td>{estudiante.Dato.Password}</td>
                        <td>{estudiante.Dato.Cursos[0]}, {estudiante.Dato.Cursos[1]}, {estudiante.Dato.Cursos[2]}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
        </>
    )
}