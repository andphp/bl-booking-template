import React, { FC, useState } from 'react'
import { useNavigate, useLocation } from 'react-router-dom'

import { Location } from 'history'
import { useLogin } from '@/api'

import styles from './index.module.less'

import Storage from '@/utils/storage'
import { LoginParams } from '@/models/login'

import bgPic from '@/assets/images/shouye.png'
import logoPic from '@/assets/images/logo.png'

const initialValues: LoginParams = {
  account_name: '',
  password: '',
  remember: true
}

const LoginPage: FC = () => {
  const loginMutation = useLogin()
  const navigate = useNavigate()
  const location = useLocation() as Location<{ from: string }>

  const token = Storage.get('accessToken')
  // const { data: currentUser, refetch } = useGetCurrentUser()
  React.useEffect(() => {
    if (token) {
      navigate('/dashboard')
    }
  }, [token])

  // 'right-panel-active'
  const [isSign, setIsSign] = useState(true)

  // const dispatch = useAppDispatch();
  const [fromValue, setFromValue] = useState(initialValues)
  const onFinished = async(fromValues: LoginParams) => {
    try {
      const result = await loginMutation.mutateAsync(fromValues)
      // console.log('result: ', result)

      if (result) {
      // localStorage.setItem('token', result.accessToken)
        Storage.set('accessToken', result.accessToken, 60 * 60 * 24 * 2)
        const from = location.state?.from || { pathname: '/dashboard/welcome' }
        // refetch()
        // setUser({ ...user, username: currentUser?.fullName || currentUser?.nickName || currentUser?.accountName || '', logged: true })

        navigate(from)
      } else {
        setTimeout(() => {
          setIsSign((state) => {
            return !state
          })
        }, 1000)
      }
    } catch (err) {
      console.log('---', err)
    }
  }

  const handleButton = (e: any) => {
    e.preventDefault()
    setIsSign(!isSign)
    onFinished(fromValue)
  }

  const getValues = (e: any) => {
    const { name, value } = e.target
    setFromValue({ ...fromValue, [name]: value })
  }

  return <div className={styles.body}>
    <div className={`${styles.container} ${isSign ? '' : styles.rightPanelActive}`}>
      {/* <!-- ?????? --> */}
      {/* <div className={`${styles.containerForm} ${styles.containerSignup}`}>
        <form className={styles.form} id='form1'>
          <h2 className={styles.formTitle}>??? ???</h2>
          <input type='text' placeholder='User' className={styles.input} />
          <input type='email' placeholder='Email' className={styles.input} />
          <input type='password' placeholder='Password' className={styles.input} />
          <button className={styles.btn}>??? ???</button>
        </form>
      </div> */}
      <div className={`${styles.containerForm} ${styles.containerSignup}`}>
        <div className={styles.formspin}>
          <div className={styles.formspinItem}>
            {/* <Spin className={styles.spin} style={{ color: '#912409' }} size='large' tip={`?????????...`} /> */}
            <span className={`${styles.antSpinDot} ${styles.antSpinDotSpin}`}>
              <i></i><i></i><i></i><i></i>
            </span>
            <p style={{ color: '#095c91' }}>
            ?????????
            </p>
          </div>
        </div>
      </div>
      {/* <!-- ?????? --> */}
      <div className={`${styles.containerForm} ${styles.containerSignin}`}>
        <form className={styles.form} id='form2'>
          <h2 className={styles.formTitle}>??? ???</h2>
          <input type='text' name='account_name' value={fromValue.account_name} onChange={getValues} placeholder='??????????????????' className={styles.input} />
          <input type='password' name='password' autoComplete='true' value={fromValue.password} onChange={getValues} placeholder='???????????????' className={styles.input} />
          {/* <a href='#' className={styles.link}>Forgot your password?</a> */}
          <a href='#' className={styles.link}>?????????????</a>
          <button className={styles.btn} onClick={handleButton}>??? ???</button>
        </form>
      </div>

      {/* <!-- ?????? --> */}
      <div className={styles.containerOverlay}>
        <div className={styles.overlay}>
          <div className={`${styles.overlayPanel} ${styles.overlayLeft}`}>
            {/* <button className={styles.btn} id='signIn' onClick={changeBtn}>??? ???</button> */}
            <img src={logoPic} alt='logo' />
          </div>
          <div className={`${styles.overlayPanel} ${styles.overlayRight}`}>
            {/* <button className={styles.btn} id='signUp' onClick={changeBtn}>??? ???</button> */}
            <img src={logoPic} alt='logo' />
          </div>
        </div>
      </div>
    </div>
    {/* <!-- ?????? --> */}
    <div className={styles.slidershow}>
      {/* <div className={styles.slidershowImage} style={{ backgroundImage: "url('https://source.unsplash.com/Snqdjm71Y5s')" }}></div>
      <div className={styles.slidershowImage} style={{ backgroundImage: "url('https://source.unsplash.com/5APj-fzKE-k')" }}></div>
      <div className={styles.slidershowImage} style={{ backgroundImage: "url('https://source.unsplash.com/wnbBH_CGOYQ')" }}></div>
      <div className={styles.slidershowImage} style={{ backgroundImage: "url('https://source.unsplash.com/OkTfw7fXLPk')" }}></div> */}
      <div className={styles.slidershowImage} style={{ backgroundImage: `url("${bgPic}")` }}></div>
      <div className={styles.slidershowImage} style={{ backgroundImage: `url("${bgPic}")` }}></div>
      <div className={styles.slidershowImage} style={{ backgroundImage: `url("${bgPic}")` }}></div>
      <div className={styles.slidershowImage} style={{ backgroundImage: `url("${bgPic}")` }}></div>
      <div className={styles.slidershowImage} style={{ backgroundImage: `url("${bgPic}")` }}></div>
    </div>

  </div>
}

export default LoginPage
