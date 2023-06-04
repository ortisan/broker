"use client";
import { useFormik } from "formik";
import { useState } from "react";
import { v4 as uuidv4 } from "uuid";
import clsx from 'clsx'
import * as Yup from "yup";
import Link from 'next/link';
import FormGroup from 'react-bootstrap/FormGroup';
import { useAuth } from '../../core/context'
import { UserModelSimple } from "../../core/models";
import { redirect } from 'next/navigation';

interface LoginResult {
  data: {
    id: string;
    email: string;
    oauth_token: string;
  };
}

interface UserInfoResult {
  data: {
    id: string;
    email: string;
    oauth_token: string;
  };
}

async function login(email: string, password: string): Promise<LoginResult> {
  return await { data: { id: uuidv4(), email: email, oauth_token: uuidv4() } };
}

async function getUserByToken(oauth_token: string): Promise<UserModelSimple> {
  return await {
    id: uuidv4(), name: "marcelo"
  };
}

export default function Login() {

  const [loading, setLoading] = useState(false)
  const { saveAuth, setCurrentUser } = useAuth()

  const loginSchema = Yup.object().shape({
    email: Yup.string()
      .email("Wrong email format")
      .min(3, "Minimum 3 symbols")
      .max(50, "Maximum 50 symbols")
      .required("Email is required"),
    password: Yup.string()
      .min(3, "Minimum 3 symbols")
      .max(50, "Maximum 50 symbols")
      .required("Password is required"),
  });

  const initialValues = {
    email: "admin@demo.com",
    password: "demo",
  };

  const formik = useFormik({
    initialValues,
    validationSchema: loginSchema,
    onSubmit: async (values, { setStatus, setSubmitting }) => {
      setLoading(true);
      try {
        const { data: auth } = await login(values.email, values.password);
        console.info(auth);
        saveAuth(auth);
        const user = await getUserByToken(auth.oauth_token);
        console.info(user);
        setCurrentUser(user);
        return redirect('/');
      } catch (error) {
        console.error(error);
        saveAuth(undefined);
        setStatus("The login details are incorrect");
        setSubmitting(false);
        setLoading(false);
      }
    },
  });

  return (
    <>
      <form
        className='form w-100'
        onSubmit={formik.handleSubmit}
        noValidate
        id='login_form'
      >

        <div className='text-center mb-11'>
          <h1 className='text-dark fw-bolder mb-3'>Sign In</h1>
          <div className='text-gray-500 fw-semibold fs-6'>Your Social Campaigns</div>
        </div>

        {/* begin::Form group */}
        <FormGroup>
          <label className='form-label fs-6 fw-bolder text-dark'>Email</label>
          <input
            placeholder='Email'
            {...formik.getFieldProps('email')}
            className={clsx(
              'form-control bg-transparent',
              { 'is-invalid': formik.touched.email && formik.errors.email },
              {
                'is-valid': formik.touched.email && !formik.errors.email,
              }
            )}
            type='email'
            name='email'
            autoComplete='off'
          />
          {formik.touched.email && formik.errors.email && (
            <div className='fv-plugins-message-container'>
              <span role='alert'>{formik.errors.email}</span>
            </div>
          )}
        </FormGroup>
        {/* end::Form group */}

        {/* begin::Form group */}
        <FormGroup>
          <label className='form-label fw-bolder text-dark fs-6 mb-0'>Password</label>
          <input
            type='password'
            autoComplete='off'
            {...formik.getFieldProps('password')}
            className={clsx(
              'form-control bg-transparent',
              {
                'is-invalid': formik.touched.password && formik.errors.password,
              },
              {
                'is-valid': formik.touched.password && !formik.errors.password,
              }
            )}
          />
          {formik.touched.password && formik.errors.password && (
            <div className='fv-plugins-message-container'>
              <div className='fv-help-block'>
                <span role='alert'>{formik.errors.password}</span>
              </div>
            </div>
          )}
        </FormGroup>
        {/* end::Form group */}

        {/* begin::Action */}
        <div className='d-grid mb-10'>
          <button
            type='submit'
            id='kt_sign_in_submit'
            className='btn btn-primary'
            disabled={formik.isSubmitting || !formik.isValid}
          >
            {!loading && <span className='indicator-label'>Continue</span>}
            {loading && (
              <span className='indicator-progress' style={{ display: 'block' }}>
                Please wait...
                <span className='spinner-border spinner-border-sm align-middle ms-2'></span>
              </span>
            )}
          </button>
        </div>
        {/* end::Action */}

        <div className='text-gray-500 text-center fw-semibold fs-6'>
          Not a Member yet?{' '}
          <Link href='/auth/registration' className='link-primary'>
            Sign up
          </Link>
        </div>

      </form>

    </>
  );
}
