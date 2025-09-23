'use server';

import { redirect } from "next/navigation";

export type FormState = {
  message: string;
};

export type AccessToken = {
  token: string
};

export async function login(state: FormState, formData: FormData) {
  const email = formData.get('email');
  const password = formData.get('password');

  try {
    const response = await fetch('http://localhost:8080/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password }),
    });

    if (!response.ok) {
      throw new Error('Login failed');
    }

    const data: AccessToken = await response.json();
    console.log(data);
    redirect('/dashboard');
  } catch (error) {
    return { message: 'Login failed' };
  }

  return { message: 'Login successful!' };
}