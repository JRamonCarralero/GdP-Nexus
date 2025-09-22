'use server';

// Define el tipo de dato que va a manejar el estado del formulario
export type FormState = {
  message: string;
};

export type AccessToken = {
  token: string
};

// La función login ahora recibe el estado anterior y el formData
// Asegúrate de que los tipos coincidan
export async function login(state: FormState, formData: FormData) {
  const email = formData.get('email');
  const password = formData.get('password');

  console.log('Login', { email, password });

  // Simula un retardo
  await new Promise((resolve) => setTimeout(resolve, 1000));

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
  } catch (error) {
    return { message: 'Login failed' };
  }

  // Devuelve un objeto del mismo tipo que FormState
  return { message: 'Login successful!' };
}