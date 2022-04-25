import React, { FC, useState, useEffect } from 'react'
// import Overview from './overview'
// import SalePercent from './salePercent'
// import TimeLine from './timeLine'
import './index.less'

const WelcomePage: FC = () => {
  const [_, setLoading] = useState(true)

  useEffect(() => {
    const timer = setTimeout(() => {
      setLoading(undefined as any)
    }, 2000)
    return () => {
      clearTimeout(timer)
    }
  }, [])
  return (
    <div>
      {/* <Overview loading={loading} />

      <SalePercent loading={loading} />
      <TimeLine loading={loading} /> */}
      welcome
    </div>
  )
}

export default WelcomePage
