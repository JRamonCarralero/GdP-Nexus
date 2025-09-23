'use client';

import { useFormState, useFormStatus } from 'react-dom';
import { register, FormState } from '@/app/(auth)/register/actions';
import Form from 'next/form';

function SubmitButton() {
  const { pending } = useFormStatus();

  return (
    <button type="submit" disabled={pending}>
      {pending ? 'Saving...' : 'Register'}
    </button>
  );
}

export function RegisterForm() {
  const initialState: FormState = { message: '' };

  const [state, formAction] = useFormState(register, initialState);

  return (
    <Form action={formAction}>
      <input type="email" placeholder="Email" name="email" required />
      <input type="password" placeholder="Password" name="password" required />
      <input type="text" placeholder="First Name" name="firstName" required />
      <input type="text" placeholder="Last Name" name="lastName" required />
      <input type="text" placeholder="Nickname" name="nickName" required />
      <SubmitButton />
      {state.message && <p>{state.message}</p>}
    </Form>
  );
}