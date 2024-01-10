import { Link } from "react-router-dom"
import imgB from '../assets/home.png';
import ret from '../assets/flecha-hacia-atras.png'
export const NavBar = () => {
    return(
        
        <nav className="nv-admin">
            <Link to='/'><img className="btn-img" src={imgB} alt="" /></Link>
            <Link to='/Admin'> <img className="btn-img" src={ret} alt="" /> </Link>
            <h2>{localStorage.getItem("user")}</h2>
            <ul>
                <li>
                    <Link to ='/Admin/Csv'> Cargar csv y Json </Link>
                </li>
                <li>
                    <Link to='/Admin/Reporte'> Reportes</Link>
                </li>
                <li>
                    <Link to={'/Admin/CBooks'}> Control de libros</Link>
                </li>
                <li>
                    <Link to = '/Admin/Tstudents'>Estudiantes</Link>
                </li>
            </ul>
        </nav>)
}
