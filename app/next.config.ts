import path from 'path'

import type { NextConfig } from 'next'

const nextConfig: NextConfig = {
  sassOptions: {
    includePaths: [path.join(__dirname, 'styles')],
    prependData: `@import "variables.scss"; @import "mixins.scss";`,
  },
}

export default nextConfig
