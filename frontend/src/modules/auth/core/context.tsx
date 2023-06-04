import {
  useContext,
  createContext,
  useState,
  Dispatch,
  SetStateAction,
  ReactNode,
  FC
} from "react";
import { AuthModel, UserModelSimple } from "./models"
import { getAuth, setAuth, removeAuth, AUTH_LOCAL_STORAGE_KEY } from "./helper"



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
  const [auth, setAuth] = useState<AuthModel | undefined>(getAuth())
  const [currentUser, setCurrentUser] = useState<UserModel | undefined>()
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

  return (
    <AuthContext.Provider value={{ auth, saveAuth, currentUser, setCurrentUser, logout }}>
      {children}
    </AuthContext.Provider>
  )
}

export { AuthProvider, useAuth }