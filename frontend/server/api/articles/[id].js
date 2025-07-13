export default defineEventHandler(async (event) => {
  try {
    const articleId = getRouterParam(event, 'id')
    
    // Use internal Docker network hostname instead of localhost
    const backendUrl = process.env.NODE_ENV === 'production' 
      ? process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080'
      : 'http://backend:8080'
    
    console.log('Server proxy fetching article from:', `${backendUrl}/api/articles/${articleId}`)
    const response = await $fetch(`${backendUrl}/api/articles/${articleId}`)
    
    return response
  } catch (error) {
    console.error('Server-side API proxy error:', error)
    throw createError({
      statusCode: 404,
      statusMessage: 'Article not found'
    })
  }
})