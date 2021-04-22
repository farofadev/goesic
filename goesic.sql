-- --------------------------------------------------------
-- Servidor:                     127.0.0.1
-- Versão do servidor:           8.0.23 - MySQL Community Server - GPL
-- OS do Servidor:               Linux
-- HeidiSQL Versão:              11.1.0.6116
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Copiando estrutura para tabela goesic.mensagens
CREATE TABLE IF NOT EXISTS `mensagens` (
  `id` char(36) NOT NULL,
  `pedido_id` char(36) NOT NULL,
  `pessoa_id` char(36) DEFAULT NULL,
  `usuario_id` char(36) DEFAULT NULL,
  `mensagem` text,
  `situacao_pedido` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `tipo` varchar(32) DEFAULT NULL,
  `criado_em` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `mensagens_pedido_id_index` (`pedido_id`),
  KEY `mensagens_pessoa_id_index` (`pessoa_id`),
  KEY `mensagens_usuario_id_index` (`usuario_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Copiando dados para a tabela goesic.mensagens: ~0 rows (aproximadamente)
/*!40000 ALTER TABLE `mensagens` DISABLE KEYS */;
/*!40000 ALTER TABLE `mensagens` ENABLE KEYS */;

-- Copiando estrutura para tabela goesic.pedidos
CREATE TABLE IF NOT EXISTS `pedidos` (
  `id` char(36) NOT NULL,
  `pessoa_id` char(36) NOT NULL,
  `situacao` varchar(36) DEFAULT NULL,
  `criado_em` datetime DEFAULT NULL,
  `data_prazo` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `pedidos_pessoa_id_index` (`pessoa_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Copiando dados para a tabela goesic.pedidos: ~0 rows (aproximadamente)
/*!40000 ALTER TABLE `pedidos` DISABLE KEYS */;
/*!40000 ALTER TABLE `pedidos` ENABLE KEYS */;

-- Copiando estrutura para tabela goesic.pessoas
CREATE TABLE IF NOT EXISTS `pessoas` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `nome` varchar(256) DEFAULT '',
  `nome_fantasia` varchar(256) DEFAULT '',
  `email` varchar(191) DEFAULT '',
  `endereco` varchar(256) DEFAULT '',
  `cidade` varchar(256) DEFAULT '',
  `estado` varchar(4) DEFAULT '',
  `cep` varchar(24) DEFAULT '',
  `documento` varchar(60) DEFAULT '',
  `tipo` varchar(60) DEFAULT '',
  `telefone` varchar(60) DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `pessoas_email_index` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Copiando dados para a tabela goesic.pessoas: ~0 rows (aproximadamente)
/*!40000 ALTER TABLE `pessoas` DISABLE KEYS */;
/*!40000 ALTER TABLE `pessoas` ENABLE KEYS */;

-- Copiando estrutura para tabela goesic.usuarios
CREATE TABLE IF NOT EXISTS `usuarios` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `nome` varchar(191) DEFAULT '',
  `email` varchar(191) DEFAULT '',
  `senha` varchar(191) DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `usuarios_email_index` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Copiando dados para a tabela goesic.usuarios: ~0 rows (aproximadamente)
/*!40000 ALTER TABLE `usuarios` DISABLE KEYS */;
/*!40000 ALTER TABLE `usuarios` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
