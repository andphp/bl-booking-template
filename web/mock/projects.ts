import { MockMethod } from 'vite-plugin-mock'

const mockProjects = {
  total: 200,
  list: [
    {
      id: 1,
      name: 'Project1',
      description: 'description'
    },
    {
      id: 2,
      name: 'Project2',
      description: 'description'
    }
  ]
}

export default [
  {
    url: '/api/v1/projects',
    method: 'get',
    response: () => {
      return mockProjects
    }
  }
] as MockMethod[]
