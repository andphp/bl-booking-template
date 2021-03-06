import React, { Suspense, useMemo } from 'react'
import ReactDOM from 'react-dom'
import { QueryClient, QueryClientProvider } from 'react-query'
// import { ReactQueryDevtools } from 'react-query-devtools'
import { RecoilRoot } from 'recoil'
import axios, { AxiosContext } from './utils/request'

import './index.css'
import App from './App'
import { ErrorBoundary } from 'react-error-boundary'
import SuspendFallbackLoading from './pages/layout/suspendFallbackLoading'
import 'nprogress/nprogress.css'
import 'default-passive-events'

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 0,
      suspense: true,
      refetchOnMount: false,
      refetchOnWindowFocus: false,
      refetchOnReconnect: true,
      refetchInterval: false
    }
  }
})

const AxiosProvider = ({ children }: React.PropsWithChildren<unknown>) => {
  const axiosValue = useMemo(() => {
    return axios
  }, [])

  return <AxiosContext.Provider value={axiosValue}>{children}</AxiosContext.Provider>
}

ReactDOM.render(
  // <React.StrictMode>
  <AxiosProvider>
    <QueryClientProvider client={queryClient}>
      <RecoilRoot>
        <ErrorBoundary
          fallbackRender={({ error, resetErrorBoundary }) => (
            <div>
              There was an error! <button onClick={() => resetErrorBoundary()}>Try again</button>
              <pre style={{ whiteSpace: 'normal' }}>{error.message}</pre>
            </div>
          )}
        >
          <Suspense fallback={<SuspendFallbackLoading />}>
            <App />
            {/* ↓ 可视化开发工具 */}
            {/* <ReactQueryDevtools /> */}
          </Suspense>
        </ErrorBoundary>
      </RecoilRoot>
    </QueryClientProvider>
  </AxiosProvider>,
  // </React.StrictMode>,
  document.getElementById('root')
)
