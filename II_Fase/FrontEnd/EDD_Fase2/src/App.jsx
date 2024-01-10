import {Routes, Route } from 'react-router-dom'
import {Login} from './pages/login'
import {Admin} from './pages/Admin'
import {Tutor} from './pages/Tutor'
import {Estudiante} from './pages/Estudiante'
import { ReadCsv } from "./components/csv"
import { TablaStudents } from './components/Tstudents'
import { LibrosT } from './components/CLibros'
import { PublicacionesT } from './components/comentarioT'
import { ControlBooks } from './components/controlBooks'
import { Reporte } from './components/reporte'
import { VisualizarLibros } from './components/librosE'

import './index.css'
export function App(){
    return(
        <div className='aplicacion'>
            <Routes>
                    <Route path='/' element={<Login/>}/>
                    <Route path='/Admin' element={<Admin/>}/>
                    <Route path="/Admin/Csv" element={<ReadCsv/>}/>
                    <Route path='/Admin/TStudents' element={<TablaStudents/>}/>
                    <Route path='/Admin/CBooks' element={<ControlBooks/>}/>
                    <Route path='/Admin/Reporte' element={<Reporte/>}/>   
                    <Route path='/Tutor' element={<Tutor/>}/>
                    <Route path='/Tutor/Libros' element={<LibrosT/>}/>
                    <Route path='/Tutor/newPublicaciones' element={<PublicacionesT/>}/>
                    <Route path='/Estudiante' element={<Estudiante/>}/>
                    <Route path='/Estudiante/VerLibros' element={<VisualizarLibros/>}/>
            </Routes>
        </div>
    )
}