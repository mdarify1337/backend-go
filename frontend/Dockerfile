# Dockerfile (frontend)
# Use an official Node.js runtime as a parent image
FROM node:20

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY ./package.json ./tsconfig.json ./


# Copy the rest of the application code
# COPY ./frontend .

# Install dependencies
RUN npm install

# Build the Next.js app
# RUN npm run build

# Expose port 3000
EXPOSE 3000

# Start the Next.js app
CMD ["npm", "run", "dev"]