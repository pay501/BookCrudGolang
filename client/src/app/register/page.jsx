"use client"

import axios from "axios"
import { useState } from "react"
import Swal from "sweetalert2"
import style from "./register.module.css"

export default function RegisterPage(){
    
    const [form , setForm] = useState({
        email:"",
        password:"",
    })
    
    function inputChange (e){
        e.preventDefault()
        const {name, value} = e.target
        setForm((prevForm)=>({...prevForm, [name]:value}))
    }

    function register (e){
        e.preventDefault()

        axios.post("http://localhost:8080/register",{
            email: form.email,
            password:form.password,
        },
        { withCredentials: true }
        ).then((res)=>{
            if(res.data.message === "Register successful"){
                Swal.fire({
                    icon:"success",
                    title:"Register successful"
                })
            }
        })
    }
    return (
        <div className={style.container}>
            <div className="flex justify-center items-center ">
                <form className="flex flex-col gap-3">
                    <div>
                        <input type="text" placeholder="email" name="email" onChange={inputChange}/>
                    </div>
                    <div>
                        <input type="password" placeholder="password" name="password" onChange={inputChange}/>
                    </div>
                    <button
                        className="bg-blue-500"
                        onClick={register}
                    >
                        Submit
                    </button>
                </form>
            </div>
        </div>
    )
}