# log 10/06/16
I spent the last few weeks trying to implement basic structure to create a renderer.
I made a attempt to create a basic library but my effort weren't focused as I am still learning the math and I am still trying to grasp the basic concepts of rendering.
I am able to make basic structure made of points and make a basic projection on a canvas.
Today I will try to make a basic rasterizer.
I learned more about matrix transforms.

# log 11/06/16
I'm learning about rotation in the transform matrix. I beleive I can't really go any farter without learning about linear algebra or geometry in general. Then I will start by making a small rasterizer and move on to raytracing.

# log 15/06/16
I did a couple of things with the vetor rotation matrices. But I've learned about a issues with floating point numbers. They are hard to work with and it forces me to be smarter in the way I build my tests.

# log 16/06/16
I've decided that it would be easier to simply render the vector on screen to see if my transformations works.
I understand how rotations matrices work. The next thing should be quaternion.

# log 17/06/16
I almost finished the scratchapixel geometry lessons. I've made a small rotation animation to test my knowledge so far. I also added the Sperical coordinate functions and stuctures.

# log 18/06/16
I finished the spherical coordinate system. There is still a couple of thing to do for the spherical system, but I'm getting bored so I've switched to something else.
I refactored the geomath package to include both slices and named type. I will build benchmark test and see what is faster. There is also the question of syntax, I will go in favor of whichever code is more readable and easy to write.

# log 21/06/16
Revamped the rasterizer. I'm trying to setup a interface for rendering.

# log 23/06/16
Refactored the geomath package.

# log 24/06/16
I started learning about the determinant of a square matrix and started learning how to invert a matrix with the guass-jordan elimination process.

# log 25/06/16
I'm unsure between creating a type for a 4 by 4 Matrix or using slices. I feel like slices would add an overhead to each operation but at the same time I could create dynamic function or method for any sized matrices and vectors.

# log 27/06/16
