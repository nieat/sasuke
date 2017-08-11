var gulp = require('gulp');
var babel = require("gulp-babel");


gulp.task('babel', function() {
	gulp.src('resource/js/*.js')
		.pipe(babel())
		.pipe(gulp.dest('public/js/'))
});

gulp.task('watch', function() {
	gulp.watch('./*.js', ['babel'])
});

gulp.task('default', ['babel', 'watch']);
