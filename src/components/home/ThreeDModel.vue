<template>
  <div ref="threeContainer" class="three-container"></div>
</template>

<script setup lang="ts">
import { onMounted, ref, nextTick } from 'vue'
import * as THREE from 'three'
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader'

const threeContainer = ref<HTMLDivElement | null>(null)

onMounted(async () => {
  await nextTick() // Ensure the DOM is updated

  if (!threeContainer.value) return

  // Scene setup
  const scene = new THREE.Scene()

  // Camera setup
  const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.01, 1000)
  camera.position.z = 5

  // Renderer setup
  const renderer = new THREE.WebGLRenderer({ alpha: true, antialias: true }) // Set alpha to true for transparent background
  renderer.setSize(window.innerWidth, window.innerHeight)
  renderer.shadowMap.enabled = true
  renderer.shadowMap.type = THREE.PCFSoftShadowMap
  threeContainer.value.appendChild(renderer.domElement)

  // Add lights
  const ambientLight = new THREE.AmbientLight(0xffffff, 0.5)
  scene.add(ambientLight)

  const directionalLight = new THREE.DirectionalLight(0xffffff, 1)
  directionalLight.position.set(1, 1, 1).normalize()
  directionalLight.castShadow = true
  scene.add(directionalLight)

  // Load GLTF model
  const loader = new GLTFLoader()
  let model: THREE.Group | undefined
  let mixer: THREE.AnimationMixer | undefined
  loader.load(
    '/models/your-model.glb',
    (gltf) => {
      model = gltf.scene
      scene.add(model)

      // Setup the animation mixer
      if (gltf.animations && gltf.animations.length) {
        mixer = new THREE.AnimationMixer(model)
        gltf.animations.forEach((clip) => {
          mixer?.clipAction(clip).play()
        })
      }
    },
    undefined,
    (error) => {
      console.error('An error happened:', error)
    }
  )

  // Animation loop
  const clock = new THREE.Clock()
  const animate = () => {
    requestAnimationFrame(animate)

    // Rotate the model if it is loaded
    if (model) {
      model.rotation.y += 0.01
    }

    // Update the mixer if it exists
    if (mixer) {
      const delta = clock.getDelta()
      mixer.update(delta)
    }

    renderer.render(scene, camera)
  }

  animate()

  // Resize listener
  window.addEventListener('resize', () => {
    const width = window.innerWidth
    const height = window.innerHeight
    renderer.setSize(width, height)
    camera.aspect = width / height
    camera.updateProjectionMatrix()
  })
})
</script>

<style scoped>
.three-container {
  width: 100%;
  height: 50vh; /* Adjust height as needed */
  position: relative; /* Ensure it stays within the document flow */
  z-index: 0;
}
</style>
