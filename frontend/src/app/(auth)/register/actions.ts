'use server';

import { redirect } from "next/navigation";

export type FormState = {
  message: string;
};

export type AccessToken = {
  token: string
};

export async function register(state: FormState, formData: FormData) {
  const email = formData.get('email');
  const password = formData.get('password');
  const firstName = formData.get('firstName');
  const lastName = formData.get('lastName');
  const nickName = formData.get('nickName');

  try {
    const response = await fetch('http://localhost:8080/auth/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password, firstName, lastName, nickName }),
    });

    if (!response.ok) {
      throw new Error('Register failed');
    }

    const data: AccessToken = await response.json();
    console.log(data);
    redirect('/dashboard');
  } catch (error) {
    return { message: 'Register failed' };
  }

  return { message: 'Register successful!' };
}