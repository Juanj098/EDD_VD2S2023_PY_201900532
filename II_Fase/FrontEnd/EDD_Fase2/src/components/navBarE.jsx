import { Link } from "react-router-dom"
import imgB from '../assets/home.png';
import ret from '../assets/flecha-hacia-atras.png'
export const NavBarE = () => {
    return(
        
        <nav className="nv-admin">
            <Link to='/'><img className="btn-img" src={imgB} alt="" /></Link>
            <h2>{localStorage.getItem("user")}</h2>  
            <Link to='/Estudiante'> <img className="btn-img" src={ret} alt="" /> </Link>
        </nav>)
}
