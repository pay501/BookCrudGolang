"use client"
import axios from 'axios'
import { useState } from 'react'
import Swal from 'sweetalert2'
import style from './login.module.css'

export default function LoginPage (){
    
    const [form, setForm] = useState({
        email:"",
        password:"",
    })

    function inputOnchange(e){
        const {name, value} = e.target
        setForm((prevForm)=>({...prevForm, [name]:value}))
    }

    const loginClick =(e)=>{
        try{
            e.preventDefault();
            if(form.email=="" || form.password==""){
                Swal.fire({
                    icon: "error",
                    title: "Oops...",
                    text: "Please fill in data!",
                    footer: '<a href="#">Why do I have this issue?</a>'
                });
            }else{
                axios.post("http://localhost:8080/login",{
                    email:form.email,
                    password:form.password,
                },
                { withCredentials: true }
                ).then(res=>{
                    if(res.data.message === "Login successful"){
                        Swal.fire({
                            icon:"success",
                            title:"Login successful"
                        })
                    }
                })
            }
        }catch(err){
            console.log("Massage:", err)
        }
    }

    return(
        <div className={style.container}>
            <form className={`flex flex-col gap-3`}>
                <div>
                    <input type="text" placeholder='Email' name='email' onChange={inputOnchange}/>
                </div>
                <div>
                    <input type="password" placeholder='Password' name='password' onChange={inputOnchange}/>
                </div>
                <button
                    className={`bg-blue-500 rounded-lg py-1`}
                    onClick={loginClick}
                >
                    Login
                </button>
            </form>
        </div>
    )
}