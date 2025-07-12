export const useContent = () => {
  const config = useRuntimeConfig()

  // Get content by key
  const getContentByKey = async (key) => {
    try {
      const response = await fetch(`${config.public.apiBase}/api/content/key/${key}`)
      if (response.ok) {
        const data = await response.json()
        return data.content || ''
      }
    } catch (error) {
      console.log(`Could not fetch content for key: ${key}`)
    }
    return ''
  }

  // Get content by section
  const getContentBySection = async (section) => {
    try {
      const response = await fetch(`${config.public.apiBase}/api/content?section=${section}`)
      if (response.ok) {
        const data = await response.json()
        return data || []
      }
    } catch (error) {
      console.log(`Could not fetch content for section: ${section}`)
    }
    return []
  }

  // Get all content
  const getAllContent = async () => {
    try {
      const response = await fetch(`${config.public.apiBase}/api/content`)
      if (response.ok) {
        const data = await response.json()
        return data || []
      }
    } catch (error) {
      console.log('Could not fetch all content')
    }
    return []
  }

  // Helper to get content with fallback
  const getContent = async (key, fallback = '') => {
    const content = await getContentByKey(key)
    return content || fallback
  }

  // Helper to get title content
  const getTitle = async (key, fallback = '') => {
    try {
      const response = await fetch(`${config.public.apiBase}/api/content/key/${key}`)
      if (response.ok) {
        const data = await response.json()
        return data.title || data.content || fallback
      }
    } catch (error) {
      console.log(`Could not fetch title for key: ${key}`)
    }
    return fallback
  }

  return {
    getContentByKey,
    getContentBySection,
    getAllContent,
    getContent,
    getTitle
  }
}