'use client';

import { useFormState, useFormStatus } from 'react-dom';
// Importa el tipo FormState del archivo actions
import { login, FormState } from '@/app/(auth)/login/actions';
import Form from 'next/form';

function SubmitButton() {
  const { pending } = useFormStatus();

  return (
    <button type="submit" disabled={pending}>
      {pending ? 'Logging in...' : 'Login'}
    </button>
  );
}

export function LoginForm() {
  // Ahora el estado inicial coincide con el tipo FormState
  // El mensaje es un string vacío, no null
  const initialState: FormState = { message: '' };
  
  const [state, formAction] = useFormState(login, initialState);

  return (
    <Form action={formAction}>
      <input type="email" placeholder="Email" name="email" required />
      <input type="password" placeholder="Password" name="password" required />
      <SubmitButton />
      {/* Muestra el mensaje si existe */}
      {state.message && <p>{state.message}</p>}
    </Form>
  );
}