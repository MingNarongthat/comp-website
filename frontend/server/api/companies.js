export default defineEventHandler(async (event) => {
  try {
    // Use internal Docker network hostname instead of localhost
    const backendUrl = process.env.NODE_ENV === 'production' 
      ? process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080'
      : 'http://backend:8080'
    
    console.log('Server proxy fetching from:', `${backendUrl}/api/companies`)
    const response = await $fetch(`${backendUrl}/api/companies`)
    
    return response
  } catch (error) {
    console.error('Server-side API proxy error:', error)
    throw createError({
      statusCode: 500,
      statusMessage: 'Failed to fetch companies from backend'
    })
  }
})