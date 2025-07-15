export default defineEventHandler(async (event) => {
  try {
    const path = getRouterParam(event, 'path')
    
    // Use internal Docker network hostname instead of localhost
    const backendUrl = process.env.NODE_ENV === 'production' 
      ? 'http://backend:8080'
      : 'http://backend:8080'
    
    const imageUrl = `${backendUrl}/uploads/${path}`
    console.log('Image proxy fetching:', imageUrl)
    
    // Fetch the image from backend
    const response = await fetch(imageUrl)
    
    if (!response.ok) {
      throw createError({
        statusCode: response.status,
        statusMessage: 'Image not found'
      })
    }
    
    // Get the image data and content type
    const imageBuffer = await response.arrayBuffer()
    const contentType = response.headers.get('content-type') || 'image/png'
    
    // Set proper headers
    setHeader(event, 'content-type', contentType)
    setHeader(event, 'cache-control', 'public, max-age=31536000') // Cache for 1 year
    
    return new Uint8Array(imageBuffer)
  } catch (error) {
    console.error('Image proxy error:', error)
    throw createError({
      statusCode: 404,
      statusMessage: 'Image not found'
    })
  }
})
