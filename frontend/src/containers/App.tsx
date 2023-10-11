import { TestID } from '@/resources/TestID'
import { getTodos } from '@/utils/api'
import { useEffect } from 'react'

export const App: React.FC = () => {
  const todos = async () => {
    console.log(await getTodos())
  }
  useEffect(() => {
    void todos()
  })
  return (
    <div id="app" data-testid={TestID.APP}>
      <h1>Next template</h1>
      <p>This is a template repository of next project.</p>
    </div>
  )
}
