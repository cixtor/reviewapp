<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Herban Legends Reviews</title>
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700|Material+Icons">
    <link rel="stylesheet" href="https://unpkg.com/bootstrap-material-design@4.0.0-beta.4/dist/css/bootstrap-material-design.min.css" integrity="sha384-R80DC0KVBO4GSTw+wZ5x2zn2pu4POSErBkf8/fSFhPXHxvHJydT0CSgAP2Yo2r4I" crossorigin="anonymous">
    <link rel="stylesheet" type="text/css" href="/assets/styles.css?v={{ .Time }}">
</head>

<body>
    <div id="reviewapp" class="container-fluid">
        <div class="row">
            <div class="col-6">
                <div class="review-card">
                    <div class="review-head" style="background-image: url('/assets/black-cherry.jpg')"></div>
                    <div class="review-body">
                        <div class="review-inside">
                            <h3>Black Cherrry Soda By Origins</h3>
                            <h6>Where, when and who</h6>
                            <p>
                                Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
                                tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
                                quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
                                consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
                                cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
                                proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
                            </p>
                            <button type="button" class="btn btn-primary" v-on:click="showReviewForm">Review Product</button>
                        </div>
                    </div>
                </div>
            </div>

            <div class="col-6">
                <!-- This could be a Vue Component, but I'm running out of time -->
                <div id="reviewform" class="section d-none">
                    <h4 class="section-title">
                        <span>Product Reviews</span>
                        <a href="#" class="section-close" v-on:click="hideReviewForm">&times;</a>
                    </h4>

                    <div id="reviewform-alert" class="alert alert-info d-none">
                        <p><!-- HTTP response --></p>
                    </div>

                    <form id="reviewform-form" @submit.prevent="submitReview">
                        <div class="form-group">
                            <label>Email Address</label>
                            <input type="email" name="email" class="form-control" placeholder="Enter email">
                            <small class="form-text text-muted">We will use Gravatar.com to set your avatar</small>
                        </div>

                        <div class="form-group">
                            <label>Name</label>
                            <input type="text" name="name" class="form-control" placeholder="Name">
                        </div>

                        <div class="form-group">
                            <label>Product Rating</label>
                            <select name="rating" class="form-control">
                                <option value="10">&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9733; &mdash; Excellent</option>
                                <option  value="9">&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9734; &mdash; Very Good</option>
                                <option  value="8">&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9734;&#9734; &mdash; Good</option>
                                <option  value="7">&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9734;&#9734;&#9734; &mdash; […]</option>
                                <option  value="6">&#9733;&#9733;&#9733;&#9733;&#9733;&#9733;&#9734;&#9734;&#9734;&#9734; &mdash; So-So (+)</option>
                                <option  value="5">&#9733;&#9733;&#9733;&#9733;&#9733;&#9734;&#9734;&#9734;&#9734;&#9734; &mdash; So-So (-)</option>
                                <option  value="4">&#9733;&#9733;&#9733;&#9733;&#9734;&#9734;&#9734;&#9734;&#9734;&#9734; &mdash; […]</option>
                                <option  value="3">&#9733;&#9733;&#9733;&#9734;&#9734;&#9734;&#9734;&#9734;&#9734;&#9734; &mdash; Bad</option>
                                <option  value="2">&#9733;&#9733;&#9734;&#9734;&#9734;&#9734;&#9734;&#9734;&#9734;&#9734; &mdash; Very Bad</option>
                                <option  value="1">&#9733;&#9734;&#9734;&#9734;&#9734;&#9734;&#9734;&#9734;&#9734;&#9734; &mdash; Horrible</option>
                            </select>
                        </div>

                        <div class="form-group">
                            <label>Comment</label>
                            <textarea name="comment" class="form-control" placeholder="What do you think about the product?"></textarea>
                        </div>

                        <button type="submit" class="btn btn-secondary">Submit Review</button>
                    </form>
                </div>

                <div id="reviewlist" class="section">
                    <h4 class="section-title">Product Reviews</h4>

                    <div class="reviews">
                        <ul>
                            <li v-for="(review,index) in reviews">
                                <div class="review-comment clearfix">
                                    <div class="review-avatar float-left">
                                        <img v-bind:src="review.avatar" />
                                    </div>

                                    <div class="review-content float-right">
                                        <div class="clearfix">
                                            <div class="review-origin float-left">
                                                <div class="review-author">${ review.name }</div>
                                                <div class="review-date">${ review.shortDate }</div>
                                            </div>

                                            <div class="review-stars float-right">
                                                <span class="review-stars-image"></span>
                                                <span v-bind:class="[review.cssScore, review.cssScoreNum]">
                                            </div>
                                        </div>

                                        <div class="review-text">
                                            <p>${ review.comment }</p>
                                        </div>
                                    </div>
                                </div>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <script type="text/javascript" src="/assets/vue.min.js"></script>
    <script type="text/javascript" src="/assets/app.js?v={{ .Time }}"></script>
</body>
</html>