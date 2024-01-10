import { Link } from "react-router-dom"
import logout from "../assets/home.png"
import ret from '../assets/flecha-hacia-atras.png'

export function NavBarT(){
    return(
        <>
            <nav className="nv-tutor">
                <Link to="/"><img className="btn-img" src={logout} alt="" /></Link>
                <Link to="/Tutor"> <img  className="btn-img" src={ret} alt="" /></Link>
                <h2>{localStorage.getItem("user")}</h2>
                <ul>
                    <li>
                        <Link to='/Tutor/newPublicaciones'> Publicaciones</Link>
                    </li>
                    <li>
                        <Link to='/Tutor/Libros'> Libros </Link>
                    </li>
                </ul>
            </nav>
        </>
    )
}