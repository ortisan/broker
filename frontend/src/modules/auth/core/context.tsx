"use client";
import {
  useContext,
  createContext,
  useState,
  Dispatch,
  SetStateAction,
  ReactNode,
  FC,
  useEffect
} from "react";
import { AuthModel, UserModel, UserModelSimple } from "./models"
import { getAuth, setAuth, removeAuth, AUTH_LOCAL_STORAGE_KEY } from "./helper"
import { useRouter } from 'next/navigation';
import { Button } from "react-bootstrap";

type AuthContextProps = {
  auth: AuthModel | undefined
  saveAuth: (auth: AuthModel | undefined) => void
  currentUser: UserModelSimple | undefined
  setCurrentUser: Dispatch<SetStateAction<UserModelSimple | undefined>>
  logout: () => void
}

const initAuthContextPropsState = {
  auth: getAuth(),
  saveAuth: () => { },
  currentUser: undefined,
  setCurrentUser: () => { },
  logout: () => { },
}

const AuthContext = createContext<AuthContextProps>(initAuthContextPropsState)

const useAuth = () => {
  return useContext(AuthContext)
}

type WithChildren = {
  children?: ReactNode
}

const AuthProvider: FC<WithChildren> = ({ children }) => {
  const router = useRouter();
  const [auth, setAuth] = useState<AuthModel | undefined>(getAuth())
  const [currentUser, setCurrentUser] = useState<UserModelSimple | undefined>()
  const saveAuth = (auth: AuthModel | undefined) => {
    setAuth(auth)
    if (auth) {
      setAuth(auth)
    } else {
      removeAuth()
    }
  }

  const logout = () => {
    saveAuth(undefined)
    setCurrentUser(undefined)
  }

  useEffect(() => {
    authCheck("teste");
  }, []);

  const authCheck = (url: string) => {
    const publicPaths = ['/login'];
    const path = url.split('?')[0];
    if (!currentUser && !publicPaths.includes(path)) {
      router.push('/login');
    }
  }

  return (
    <AuthContext.Provider value={{ auth, saveAuth, currentUser, setCurrentUser, logout }}>
      <div>Id: {currentUser?.id}</div>
      <div>Name: {currentUser?.name}</div>
      <Button onClick={logout}>Logout</Button>
      {children}
    </AuthContext.Provider>
  )
}

export { AuthProvider, useAuth }