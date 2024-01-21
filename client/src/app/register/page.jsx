"use client"

import axios from "axios"
import { useRouter } from "next/navigation"
import { useState } from "react"
import Swal from "sweetalert2"
import style from "./register.module.css"

export default function RegisterPage(){
    
    const router = useRouter()
    const [form , setForm] = useState({
        email:"",
        password:"",
    })
    
    function inputOnchange (e){
        e.preventDefault()
        const {name, value} = e.target
        setForm((prevForm)=>({...prevForm, [name]:value}))
    }

    function register (e){
        e.preventDefault()
        if( form.email ==="" || form.password ===""){
            Swal.fire({
                icon:'error',
                title:"Oops! Register failed!",
                text:"Please, fill in all data."
            })
        }else{
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
                    router.push("/")
                }else{
                    Swal.fire({
                        icon:"warning",
                        title:"Email has been used"
                    })
                }
            })
        }
    }

    console.log(form)
    return (
        <div className={style.container}>
            <div className="hero min-h-screen bg-base-200">
                <div className="hero-content flex gap-10">
                    <div className="text-center flex-1">
                        <h1 className="text-5xl font-bold">Register now!</h1>
                        <p className="py-6">Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a id nisi.</p>
                    </div>
                    <div className="card  max-w-lg shadow-2xl bg-base-100 flex-1">
                        <form className="card-body">
                            <div className="form-control">
                                <label className="label">
                                    <span className="label-text">Email</span>
                                </label>
                                <input
                                    type="email"
                                    placeholder="email"
                                    className="input input-bordered"
                                    required
                                    name='email'
                                    onChange={inputOnchange}
                                />
                            </div>
                            <div className="form-control">
                                <label className="label">
                                    <span className="label-text">Password</span>
                                </label>
                                <input
                                    type="password"
                                    placeholder="password"
                                    className="input input-bordered"
                                    required
                                    name='password'
                                    onChange={inputOnchange}
                                />
                            </div>
                            <div className="form-control mt-6">
                                <button
                                    className="btn btn-primary"
                                    onClick={register}
                                >
                                    Submit
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    )
}