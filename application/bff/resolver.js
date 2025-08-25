export const resolvers = {
  Query: {
    book: async (parent, args, context) => {
      // const response = await context.dataSources.catalogueApi.getBook(args.id)
      const response = await context.dataSources.catalogueApi.getBook(args.id)
      return response
    },
    books: async (parent, args, context) => {
      // const response = await context.dataSources.catalogueApi.listBooks()
      const response = await context.dataSources.catalogueApi.listBooks()
      return response
    }
  }
}