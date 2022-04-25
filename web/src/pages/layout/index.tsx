import { useGetCurrentMenus, useGetCurrentUser } from '@/api'
import { useLocale } from '@/locales'
import { userState } from '@/stores/user'
import { FrownOutlined, HeartOutlined, MenuFoldOutlined, MenuUnfoldOutlined, SmileOutlined } from '@ant-design/icons'
import type { MenuDataItem } from '@ant-design/pro-layout'
import ProLayout from '@ant-design/pro-layout'
import { createBrowserHistory } from 'history'
import React, { FC, useEffect, useState } from 'react'
import { Link, useLocation, useNavigate } from 'react-router-dom'
import { useRecoilState } from 'recoil'
import { useGuide } from '../guide/useGuide'
import Footer from './components/Footer'
import GlobalHeaderRight from './components/GlobalHeaderRight'
import TabRoute from './components/TabRoute'
import styles from './index.module.less'
import _ from 'lodash'
import Icon from '@/components/icon'
import logoPic from '@/assets/images/logo_min.png'
import { Logged } from '@/utils/helper'

const history = createBrowserHistory()

const IconMap: { [key: string]: React.ReactNode } = {
  smile: <SmileOutlined />,
  heart: <HeartOutlined />,
  frown: <FrownOutlined />
}

const LayoutPage: FC = () => {
  const { data: menuList, refetch: menuRefetch } = useGetCurrentMenus()
  const { data: currentUser, refetch: userRefetch } = useGetCurrentUser()
  const [user, setUser] = useRecoilState(userState)

  const { device, collapsed, newUser, settings, logged } = user
  const isMobile = device === 'MOBILE'
  const { driverStart } = useGuide()
  const location = useLocation()
  const navigate = useNavigate()
  const { formatMessage } = useLocale()
  useEffect(() => {
    if (!Logged()) {
      navigate('/login', { replace: true })
    }
    if (location.pathname === '/') {
      navigate('/dashboard')
    }
  }, [navigate, location])

  const toggle = () => {
    setUser({ ...user, collapsed: !collapsed })
  }
  useEffect(() => {
    if (!menuList) {
      menuRefetch()
    }
  }, [menuList])

  useEffect(() => {
    if (currentUser) {
      setUser({ ...user, username: currentUser?.fullName || currentUser?.nickName || currentUser?.accountName || '', logged: true })
    } else {
      userRefetch()
    }
  }, [currentUser])

  useEffect(() => {
    if (Logged() && logged && newUser) {
      driverStart()
    }
  }, [logged])

  const loopMenuItem = (menus?: MenuDataItem[]): MenuDataItem[] => {
    if (!menus) return []

    const m = menus.map(({ icon, children, ...item }) => ({
      ...item,
      hideInBreadcrumb: false,
      icon: icon && <Icon type={`icon-${icon}`} />,
      children: children && loopMenuItem(children)
    }))

    return m
  }
  const [showMaximize, setShowMaximize] = useState(true)

  const clickChangeMaximize = () => {
    setShowMaximize(!showMaximize)
  }
  const contentStyle = showMaximize
    ? { margin: 0 }
    : {
      margin: 0,
      position: 'fixed',
      left: 0,
      width: '100%',
      zIndex: 100,
      minHeight: '100vh',
      backgroundColor: '#f0f2f5'
    }
  const ProLayoutElement = (
    <ProLayout
      fixSiderbar
      collapsed={collapsed}
      location={{
        pathname: location.pathname
      }}
      contentStyle={contentStyle}
      {...settings}
      onCollapse={toggle}
      formatMessage={formatMessage}
      onMenuHeaderClick={() => history.push('/')}
      logo={logoPic}
      title='BOOKING'
      headerTitleRender={(logo, title, props) => (
        <a className={styles.layoutPageHeader}>
          {logo}
          {title}
        </a>
      )}
      headerHeight={58}
      // menuHeaderRender={undefined}
      menuItemRender={(menuItemProps, defaultDom) => {
        if (menuItemProps.isUrl || !menuItemProps.path || location.pathname === menuItemProps.path) {
          return defaultDom
        }

        return <Link to={menuItemProps.path}>{defaultDom}</Link>
      }}
      breadcrumbRender={(routers = []) => [
        {
          path: '/',
          breadcrumbName: formatMessage({ id: 'menu.home' })
        },
        ...routers
      ]}
      itemRender={(route, params, routes, paths) => {
        const first = routes.indexOf(route) === 0
        return first ? <Link to={paths.join('/')}>{route.breadcrumbName}</Link> : <span>{route.breadcrumbName}</span>
      }}
      menuDataRender={() => loopMenuItem(menuList)}
      // menuDataRender={() => m}
      rightContentRender={() => <GlobalHeaderRight />}
      footerRender={() => <Footer />}
      collapsedButtonRender={() => {
        return (
          <div
            onClick={() => toggle}
            style={{
              cursor: 'pointer',
              fontSize: '16px'
            }}
          >
            <span id='sidebar-trigger'>{collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}</span>
          </div>
        )
      }}
    >
      {TabRoute(clickChangeMaximize, showMaximize)}
    </ProLayout>
  )

  return <>{ProLayoutElement}</>
}

export default LayoutPage
