-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 25, 2023 at 04:47 PM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `cvwoforum`
--

-- --------------------------------------------------------

--
-- Table structure for table `comments`
--

CREATE TABLE `comments` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `content` longtext DEFAULT NULL,
  `post_id` bigint(20) UNSIGNED DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `creator` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `comments`
--

INSERT INTO `comments` (`id`, `content`, `post_id`, `user_id`, `created_at`, `creator`) VALUES
(27, 'I agree with you', 19, 1, '2023-01-25 23:39:52.684', 'NUSadmin'),
(28, 'I agree with myself too.', 19, 3, '2023-01-25 23:40:01.600', 'NewAccount'),
(29, 'Who else likes bread?', 19, 3, '2023-01-25 23:40:19.341', 'NewAccount'),
(30, 'I don\'t agree, bread is dry', 19, 2, '2023-01-25 23:41:07.939', 'Ha Thu'),
(31, 'Super idol 的笑容都没你的甜\n八月正午的阳光都没你耀眼\n热爱105度的你 滴滴清纯的蒸馏水', 19, 1, '2023-01-25 23:41:27.930', 'NUSadmin'),
(32, 'HELP', 20, 2, '2023-01-25 23:46:01.124', 'Ha Thu'),
(33, 'The more we get together\nTogether, together\nThe more we get together\nThe happier we\'ll be\n\'Cause your friends are my friends\nAnd my friends are your friends\nThe more we get together\nThe happier we\'ll be\nOh, the more we get together\nTogether, together\nThe more we get together\nThe happier we\'ll be\nThere\'s Chris and Tanya\nAnd Jason and Jusitn\nThe more we get together\nThe happier we\'ll be\nThe more we get together\nTogether, together\nThe more we get together\nThe happier we\'ll be\n\'Cause your friends are my friends\nAnd my friends are your friends\nThe more we get together\nThe happier we\'ll be', 3, 1, '2023-01-25 23:46:41.144', 'NUSadmin'),
(34, 'The more we get together\nTogether, together\nThe more we get together\nThe happier we\'ll be\n\'Cause your friends are my friends\nAnd my friends are your friends\nThe more we get together\nThe happier we\'ll be\nOh, the more we get together\nTogether, together\nThe more we get together\nThe happier we\'ll be\nThere\'s Chris and Tanya\nAnd Jason and Jusitn\nThe more we get together\nThe happier we\'ll be\nThe more we get together\nTogether, together\nThe more we get together\nThe happier we\'ll be\n\'Cause your friends are my friends\nAnd my friends are your friends\nThe more we get together\nThe happier we\'ll be', 19, 3, '2023-01-25 23:46:53.807', 'NewAccount'),
(35, 'YTA - the older cat was there first and your roommate warned you that he did not like other cats. Instead of having any kind of conversation, you just brought home a kitten and now expect an elderly cat to be confined to one room for the rest of his life. You are most indeed the assholes here. That cat deserves to live out his life in peace and if you want to keep your kitten, you better go find somewhere else to live. Your roommate has every right to be angry and you are out of line.\n\nEDIT - OP, I was so delighted to see your edits this a.m.!! You are now on the path to becoming an excellent cat papa (sorry!) mama, something this world always needs more of! I\'m so glad your RM\'s kitty was willing to interact with you and while that kitty may never become a true \"grandparent\" cat, at least now there is hope that the two felines can live together in peace. Bravo and well done!\n\nAnd thank you to everyone for the awards!!!', 10, 2, '2023-01-25 23:47:17.433', 'Ha Thu'),
(36, 'YTA. First, the cat is not \"mean\" just because it does not like to be picked up or petted by you. It\'s a pet with it\'s own preferences, which is completely normal. Maybe it would like you more if you did not do the things he dislikes? Second, you brought a cat without talking to your roommate and now demand her cat should stay in her room? Her cat is elderly and has been living there already, it\'s basically his flat, and to expect an elderly pet to be confined to one room is cruel, selfish and unfair for him. You also sound selfish for not understanding cat behaviour and just bringing one on a whim without even talking or trying out if they get along with each other. That\'s just dumb and basic pet owner knowledge. Many shelters won\'t even let you adopt without knowing they will get along. If you really want to keep your cat, you should consider moving out to avoid stress for both of them (and yourself and roommate).', 10, 3, '2023-01-25 23:47:31.281', 'NewAccount');

-- --------------------------------------------------------

--
-- Table structure for table `communities`
--

CREATE TABLE `communities` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext DEFAULT NULL,
  `desc` longtext DEFAULT NULL,
  `image_url` longtext DEFAULT NULL,
  `link` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `communities`
--

INSERT INTO `communities` (`id`, `name`, `desc`, `image_url`, `link`) VALUES
(1, 'NUS', 'To facilitate discussions on NUS', 'https://cde.nus.edu.sg/wp-content/uploads/2021/08/article-image-2048x1152-1.png', 'nus'),
(2, 'CVWO', 'To facilitate discussions on the web development internship CVWO', 'https://www.comp.nus.edu.sg/images/resources/20220210_COM06677_ed.jpg', 'cvwo'),
(3, 'Cats', 'Any discussion regarding cats is allowed. Discussion of furries or the Cats movie is not allowed.', 'https://i.guim.co.uk/img/media/26392d05302e02f7bf4eb143bb84c8097d09144b/446_167_3683_2210/master/3683.jpg?width=1200&quality=85&auto=format&fit=max&s=a52bbe202f57ac0f5ff7f47166906403', 'cats'),
(8, 'Baking', 'For disscussions about anything related to baking.', 'https://www.theperfectloaf.com/wp-content/uploads/2015/12/theperfectloaf-mybestsourdoughrecipe-title-1.jpg', 'baking');

-- --------------------------------------------------------

--
-- Table structure for table `posts`
--

CREATE TABLE `posts` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `desc` longtext DEFAULT NULL,
  `content` longtext DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `community_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `posts`
--

INSERT INTO `posts` (`id`, `created_at`, `title`, `desc`, `content`, `user_id`, `community_id`) VALUES
(3, '2023-01-24 21:20:08.958', 'Task Manager/Daily Planner with Canvas Integration', 'I came up with a task manager/ daily planner with Canvas Integration! Hope you find it helpful. \nOP: https://www.reddit.com/r/nus/comments/10jzcxo/task_managerdaily_planner_with_canvas_integration/', 'Do you sometimes find it difficult to keep track of all your deadlines? Do you want an easier way to stay up to date with the latest information available on Canvas? Are you looking for an app to help you stay organised and be more productive this new year?\n\nIf any of these questions relate to you, allow me to make a case for you to try out Yet Another Todo App!\n\nYata is a task manager and daily planner available on iOS, which started out as a side project because my answer to all of the above questions was a resounding yes.\n\nFirstly, it is completely free: no IAPs or subscriptions, and no sign-up required!\n\nSecondly, Yata integrates with Canvas, allowing you to easily consolidate your modules’ assignments, quizzes and their respective deadlines with a single tap!\n\nIt is able to retrieve the necessary information such as a quiz’s title and deadline, which can be immediately saved as tasks in the app.\n\nLastly, Yata on its own is a capable task manager and daily planner. It has many core features that similar apps have, such as natural language recognition, iCloud sync, iOS calendar integration, and much more.\n\nIf you’re interested, do check it out on the App Store for more details (again, it is completely free)!\n\nHello all', 1, 1),
(10, '2023-01-25 00:32:49.707', 'AITA for asking my roommate to keep her cat in her bedroom?', 'Please help me judge if I am the asshole.', 'OP: https://www.reddit.com/r/AmItheAsshole/comments/10jot1f/aita_for_asking_my_roommate_to_keep_her_cat_in/\n\nI (24F) live with my girlfriend (25F) and a roommate (23F). We share a two bedroom apartment and have lived together for 2 years now. Up until now, we have gotten along pretty well.\n\nMy roommate has an elderly cat. The cat is honestly kinda mean, she doesn’t like me or my girlfriend and will run away if we try to pet her or pick her up. So we just avoid each other.\n\nMy roommate asked if we had cats prior to moving in because she said her cat “doesn’t like other cats.” We did not have any pets at the time so it worked out. We did tell her that we might want one someday to which she said we would have to revisit the issue if it came up.\n\nWell, my girlfriend and I have gotten to a place where we really want a pet of our own. We saw a kitten at the shelter and just fell in love with him. We got the okay from our landlord and brought him home last week! His name is Banana Pudding! :)\n\nOur roommate is furious with us. She told us that we were aware that her cat doesn’t like other cats. I told her honestly that it’s unfair for us to have to tiptoe around her cat and that we were allowed to have our own pets as long as the landlord is okay with it (which he is). She couldn’t expect us to never get a pet or to cater to her cat’s needs 24/7. The conversation ended there.\n\nUnfortunately things have escalated because her cat hisses and swats at Banana Pudding just for existing. She is honestly pretty aggressive which is a big issue imo. Our kitten is very friendly and sweet and causes no problems.\n\nWe asked our roommate that she keep her cat in her bedroom since she’s elderly and aggressive. She refused and hasn’t been speaking to us. She says we are huge AHs and even got some mutual friends to take her side. I personally think that if her cat is so aggressive and uncomfortable, it’s her responsibility to live alone. It’s gotten so bad that my girlfriend is crying and considering taking the kitten back to the shelter. So here I am to get an outside perspective because idk what to think now.\n\nAITA?\n\nEDIT: So I am definitely the AH. The comments have really helped me to understand that my roommate’s cat is just acting… well, like a cat! I am willing to admit I was ignorant and impulsive. I have never had a pet before and should have done more research before jumping into pet ownership. I plan to sincerely apologize to my roommate. She is willing to talk once she gets home from work. I am going to see if she’s willing to try a proper introduction for both kitties and Banana Pudding will be staying in our room for the duration while her kitty will be able to roam. I am hopeful that things will work out and I want to thank everyone for being blunt and educating me. I was impulsive and inconsiderate, but my roommate is a great person and I think she will forgive us and be willing to try to make it work.\n\nEDIT #2: Roommate and I have talked. I apologized to her for being selfish, ignorant, and for putting her cat in an uncomfortable situation. She forgave me and said she’d be happy to help do a proper introduction between the kitties. She said that if we do things slowly and carefully, her cat might be able to accept Banana Pudding. For now, our kitten will be staying in our room for the introduction process. She also taught me a little bit about cat body language and guys… her cat came and sat in my lap! She even purred! My roommate is a very kind and forgiving person and I’m very thankful that she chose to forgive me after I was such a huge AH to her. She played with Banana Pudding and gave me some cat food recommendations. I think things are gonna work out. Thanks everyone for your advice and criticism!', 2, 3),
(15, '2023-01-25 13:21:10.915', '3am existential crisis', 'help', 'ugh idek where to start, its like my final sem in uni, im struggling to finish my fyp, my acquaintances or following on insta are all getting married or getting bto, and meanwhile i can\'t even seem to make close, genuine friends anymore since like jc?\n\ni haven\'t been in a clique since upper sec, somehow i made it through uni without close friends just by being a straggler for every mod and begging to join open groups, sometimes taking Ls when i get not great groupmates, but overall still managing a 3.2 cap which is far from great but im happy to graduate with.\n\nbut one thing rly bothers me, and it\'s that despite many ppl saying uni is where u discover urself and find close friends, i never rly did that. i did meet a small handful of cool and fun people, but even they have their own cliques and/or partners to hang out with. i kinda regret not staying in any rc or hall, bc staying in utr didn\'t rly help in making friends (apartment mates all kinda hv their own things/people), and i don\'t get why most clubs in nus cost money to join.\n\nto any seniors who had a similar lonely experience as me, does it get better in the working world? honestly im just rly miserable bc i feel like im sorta transparent or invisible, i dun hv it as bad as many unfortunate people but i am so far from the best too, so idk how to feel abt things. plz no delete', 1, 1),
(19, '2023-01-25 23:39:21.364', 'I love bread', 'It is so damn tasty.', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 3, 8),
(20, '2023-01-25 23:45:55.855', 'I am dying, Golang is very hard', 'Help please', 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.', 2, 2);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `username` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `password` longblob DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `password`) VALUES
(1, 'NUSadmin', 'e0969496@u.nus.edu', 0x2432612431342432656a56783336616d6a767946334364594b6a37692e52616b4a3673674364586a6f4f414c52665553317534774c3551676e583147),
(2, 'Ha Thu', 'hathu2516@gmail.com', 0x243261243134244b624a4b66736c3145724c6a4d4a457676702f66542e426d4266786552706b6a45687161526b464673644a47613331466d38794271),
(3, 'NewAccount', 'hello@email.com', 0x2432612431342435775345546c6339445959367a7430624c394f7937652f307362593142747a45765363705447786e55476d31746c444c5730765275);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `comments`
--
ALTER TABLE `comments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_comments_post` (`post_id`),
  ADD KEY `fk_comments_user` (`user_id`);

--
-- Indexes for table `communities`
--
ALTER TABLE `communities`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_posts_user` (`user_id`),
  ADD KEY `fk_posts_community` (`community_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `comments`
--
ALTER TABLE `comments`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=37;

--
-- AUTO_INCREMENT for table `communities`
--
ALTER TABLE `communities`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `posts`
--
ALTER TABLE `posts`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `comments`
--
ALTER TABLE `comments`
  ADD CONSTRAINT `fk_comments_post` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`),
  ADD CONSTRAINT `fk_comments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `posts`
--
ALTER TABLE `posts`
  ADD CONSTRAINT `fk_posts_community` FOREIGN KEY (`community_id`) REFERENCES `communities` (`id`),
  ADD CONSTRAINT `fk_posts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
