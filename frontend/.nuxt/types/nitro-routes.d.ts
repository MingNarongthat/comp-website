// Generated by nitro
import type { Serialize, Simplify } from "nitropack/types";
declare module "nitropack/types" {
  type Awaited<T> = T extends PromiseLike<infer U> ? Awaited<U> : T
  interface InternalApi {
    '/api/articles': {
      'default': Simplify<Serialize<Awaited<ReturnType<typeof import('../../server/api/articles').default>>>>
    }
    '/api/articles/:id': {
      'default': Simplify<Serialize<Awaited<ReturnType<typeof import('../../server/api/articles/[id]').default>>>>
    }
    '/api/companies': {
      'default': Simplify<Serialize<Awaited<ReturnType<typeof import('../../server/api/companies').default>>>>
    }
    '/api/proxy/uploads/**:path': {
      'default': Simplify<Serialize<Awaited<ReturnType<typeof import('../../server/api/proxy/uploads/[...path]').default>>>>
    }
    '/__nuxt_error': {
      'default': Simplify<Serialize<Awaited<ReturnType<typeof import('../../node_modules/nuxt/dist/core/runtime/nitro/handlers/renderer').default>>>>
    }
    '/__nuxt_island/**': {
      'default': Simplify<Serialize<Awaited<ReturnType<typeof import('../../server/#internal/nuxt/island-renderer').default>>>>
    }
  }
}
export {}