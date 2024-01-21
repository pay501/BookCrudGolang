"use client"
import axios from 'axios'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { useState } from 'react'
import Swal from 'sweetalert2'
import style from './login.module.css'

export default function LoginPage() {

    const router = useRouter()
    const [form, setForm] = useState({
        email: "",
        password: "",
    })

    function inputOnchange(e) {
        const { name, value } = e.target
        setForm((prevForm) => ({ ...prevForm, [name]: value }))
    }
    
    const loginClick = (e) => {
        try {
            e.preventDefault();
            if (form.email == "" || form.password == "") {
                Swal.fire({
                    icon: "error",
                    title: "Oops...",
                    text: "Please fill in data!",
                    footer: '<a href="#">Why do I have this issue?</a>'
                });
            } else {
                axios.post("http://localhost:8080/login", {
                    email: form.email,
                    password: form.password,
                },
                    { withCredentials: true }
                ).then(res => {
                    if (res.data.message === "Login successful") {
                        Swal.fire({
                            icon: "success",
                            title: "Login successful"
                        })
                        router.push("/")
                    }else{
                        Swal.fire({
                            icon: "error",
                            title: "Login failed",
                            text:"Email or Password is wrong!"
                        })
                    }
                })
            }
        } catch (err) {
            console.log("Massage:", err)
        }
    }
    return (
        <div className={style.container}>
            <div className="hero min-h-screen bg-base-200">
                <div className="hero-content flex gap-10">
                    <div className="text-center flex-1">
                        <h1 className="text-5xl font-bold">Login now!</h1>
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
                                <label className="label">
                                    <a href="/" className="label-text-alt link link-hover">
                                        Forgot password?
                                    </a>
                                </label>
                                <span className='ml-1'>Haven't register yet? <Link href="/register" className='link link-hover text-blue-600'>Register!</Link></span>
                            </div>
                            <div className="form-control mt-6">
                                <button
                                    className="btn btn-primary"
                                    onClick={loginClick}
                                >
                                    Login
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    )
}