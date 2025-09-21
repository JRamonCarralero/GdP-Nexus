'use server';

// Define el tipo de dato que va a manejar el estado del formulario
export type FormState = {
  message: string;
};

// La función login ahora recibe el estado anterior y el formData
// Asegúrate de que los tipos coincidan
export async function login(state: FormState, formData: FormData) {
  const email = formData.get('email');
  const password = formData.get('password');

  console.log('Login', { email, password });

  // Simula un retardo
  await new Promise((resolve) => setTimeout(resolve, 1000));

  // Devuelve un objeto del mismo tipo que FormState
  return { message: 'Login successful!' };
}