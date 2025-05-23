PGDMP                      }         	   Gastrobar    16.8    16.8 :    `           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            a           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            b           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            c           1262    16793 	   Gastrobar    DATABASE     q   CREATE DATABASE "Gastrobar" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'es-ES';
    DROP DATABASE "Gastrobar";
                postgres    false            �            1259    16809 
   menu_items    TABLE     �  CREATE TABLE public.menu_items (
    id integer NOT NULL,
    nombre character varying(100) NOT NULL,
    categoria character varying(50) NOT NULL,
    precio numeric(10,2) NOT NULL,
    stock integer NOT NULL,
    CONSTRAINT menu_items_categoria_check CHECK (((categoria)::text = ANY ((ARRAY['comida'::character varying, 'bebida'::character varying])::text[]))),
    CONSTRAINT menu_items_stock_check CHECK ((stock >= 0))
);
    DROP TABLE public.menu_items;
       public         heap    postgres    false            �            1259    16808    menu_items_id_seq    SEQUENCE     �   CREATE SEQUENCE public.menu_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.menu_items_id_seq;
       public          postgres    false    218            d           0    0    menu_items_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.menu_items_id_seq OWNED BY public.menu_items.id;
          public          postgres    false    217            �            1259    16840    order_items    TABLE     �   CREATE TABLE public.order_items (
    id integer NOT NULL,
    order_id integer NOT NULL,
    menu_item_id integer NOT NULL,
    cantidad integer NOT NULL,
    CONSTRAINT order_items_cantidad_check CHECK ((cantidad > 0))
);
    DROP TABLE public.order_items;
       public         heap    postgres    false            �            1259    16839    order_items_id_seq    SEQUENCE     �   CREATE SEQUENCE public.order_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public.order_items_id_seq;
       public          postgres    false    224            e           0    0    order_items_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.order_items_id_seq OWNED BY public.order_items.id;
          public          postgres    false    223            �            1259    16827    orders    TABLE     �   CREATE TABLE public.orders (
    id integer NOT NULL,
    table_id integer NOT NULL,
    total numeric(10,2) NOT NULL,
    paid boolean DEFAULT false
);
    DROP TABLE public.orders;
       public         heap    postgres    false            �            1259    16826    orders_id_seq    SEQUENCE     �   CREATE SEQUENCE public.orders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.orders_id_seq;
       public          postgres    false    222            f           0    0    orders_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;
          public          postgres    false    221            �            1259    16867    permissions_logs    TABLE     �   CREATE TABLE public.permissions_logs (
    id integer NOT NULL,
    user_id integer NOT NULL,
    action character varying(255) NOT NULL,
    status character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
 $   DROP TABLE public.permissions_logs;
       public         heap    postgres    false            �            1259    16866    permissions_logs_id_seq    SEQUENCE     �   CREATE SEQUENCE public.permissions_logs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.permissions_logs_id_seq;
       public          postgres    false    228            g           0    0    permissions_logs_id_seq    SEQUENCE OWNED BY     S   ALTER SEQUENCE public.permissions_logs_id_seq OWNED BY public.permissions_logs.id;
          public          postgres    false    227            �            1259    16858    reports    TABLE     �   CREATE TABLE public.reports (
    id integer NOT NULL,
    titulo character varying(255) NOT NULL,
    descripcion text NOT NULL,
    fecha date NOT NULL
);
    DROP TABLE public.reports;
       public         heap    postgres    false            �            1259    16857    reports_id_seq    SEQUENCE     �   CREATE SEQUENCE public.reports_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.reports_id_seq;
       public          postgres    false    226            h           0    0    reports_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.reports_id_seq OWNED BY public.reports.id;
          public          postgres    false    225            �            1259    16818    tables    TABLE       CREATE TABLE public.tables (
    id integer NOT NULL,
    estado character varying(50) DEFAULT 'libre'::character varying,
    CONSTRAINT tables_estado_check CHECK (((estado)::text = ANY ((ARRAY['libre'::character varying, 'ocupada'::character varying])::text[])))
);
    DROP TABLE public.tables;
       public         heap    postgres    false            �            1259    16817    tables_id_seq    SEQUENCE     �   CREATE SEQUENCE public.tables_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.tables_id_seq;
       public          postgres    false    220            i           0    0    tables_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.tables_id_seq OWNED BY public.tables.id;
          public          postgres    false    219            �            1259    16795    users    TABLE     �  CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    username character varying(50) NOT NULL,
    password character varying(255) NOT NULL,
    role character varying(20) NOT NULL,
    CONSTRAINT users_role_check CHECK (((role)::text = ANY ((ARRAY['Dueño'::character varying, 'Admin'::character varying, 'Employee'::character varying])::text[])))
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    16794    users_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    216            j           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    215            �           2604    16812    menu_items id    DEFAULT     n   ALTER TABLE ONLY public.menu_items ALTER COLUMN id SET DEFAULT nextval('public.menu_items_id_seq'::regclass);
 <   ALTER TABLE public.menu_items ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    218    217    218            �           2604    16843    order_items id    DEFAULT     p   ALTER TABLE ONLY public.order_items ALTER COLUMN id SET DEFAULT nextval('public.order_items_id_seq'::regclass);
 =   ALTER TABLE public.order_items ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    224    223    224            �           2604    16830 	   orders id    DEFAULT     f   ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);
 8   ALTER TABLE public.orders ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    222    221    222            �           2604    16870    permissions_logs id    DEFAULT     z   ALTER TABLE ONLY public.permissions_logs ALTER COLUMN id SET DEFAULT nextval('public.permissions_logs_id_seq'::regclass);
 B   ALTER TABLE public.permissions_logs ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    228    227    228            �           2604    16861 
   reports id    DEFAULT     h   ALTER TABLE ONLY public.reports ALTER COLUMN id SET DEFAULT nextval('public.reports_id_seq'::regclass);
 9   ALTER TABLE public.reports ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    225    226    226            �           2604    16821 	   tables id    DEFAULT     f   ALTER TABLE ONLY public.tables ALTER COLUMN id SET DEFAULT nextval('public.tables_id_seq'::regclass);
 8   ALTER TABLE public.tables ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    219    220    220            �           2604    16798    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    215    216            S          0    16809 
   menu_items 
   TABLE DATA           J   COPY public.menu_items (id, nombre, categoria, precio, stock) FROM stdin;
    public          postgres    false    218   �A       Y          0    16840    order_items 
   TABLE DATA           K   COPY public.order_items (id, order_id, menu_item_id, cantidad) FROM stdin;
    public          postgres    false    224   �B       W          0    16827    orders 
   TABLE DATA           ;   COPY public.orders (id, table_id, total, paid) FROM stdin;
    public          postgres    false    222   �B       ]          0    16867    permissions_logs 
   TABLE DATA           S   COPY public.permissions_logs (id, user_id, action, status, created_at) FROM stdin;
    public          postgres    false    228   �B       [          0    16858    reports 
   TABLE DATA           A   COPY public.reports (id, titulo, descripcion, fecha) FROM stdin;
    public          postgres    false    226   ZC       U          0    16818    tables 
   TABLE DATA           ,   COPY public.tables (id, estado) FROM stdin;
    public          postgres    false    220   �C       Q          0    16795    users 
   TABLE DATA           J   COPY public.users (id, name, email, username, password, role) FROM stdin;
    public          postgres    false    216   �C       k           0    0    menu_items_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.menu_items_id_seq', 13, true);
          public          postgres    false    217            l           0    0    order_items_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public.order_items_id_seq', 1, false);
          public          postgres    false    223            m           0    0    orders_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.orders_id_seq', 1, true);
          public          postgres    false    221            n           0    0    permissions_logs_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public.permissions_logs_id_seq', 2, true);
          public          postgres    false    227            o           0    0    reports_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.reports_id_seq', 1, true);
          public          postgres    false    225            p           0    0    tables_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.tables_id_seq', 1, false);
          public          postgres    false    219            q           0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 7, true);
          public          postgres    false    215            �           2606    16816    menu_items menu_items_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.menu_items
    ADD CONSTRAINT menu_items_pkey PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.menu_items DROP CONSTRAINT menu_items_pkey;
       public            postgres    false    218            �           2606    16846    order_items order_items_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.order_items DROP CONSTRAINT order_items_pkey;
       public            postgres    false    224            �           2606    16833    orders orders_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_pkey;
       public            postgres    false    222            �           2606    16873 &   permissions_logs permissions_logs_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY public.permissions_logs
    ADD CONSTRAINT permissions_logs_pkey PRIMARY KEY (id);
 P   ALTER TABLE ONLY public.permissions_logs DROP CONSTRAINT permissions_logs_pkey;
       public            postgres    false    228            �           2606    16865    reports reports_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.reports
    ADD CONSTRAINT reports_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.reports DROP CONSTRAINT reports_pkey;
       public            postgres    false    226            �           2606    16825    tables tables_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.tables
    ADD CONSTRAINT tables_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.tables DROP CONSTRAINT tables_pkey;
       public            postgres    false    220            �           2606    16805    users users_email_key 
   CONSTRAINT     Q   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);
 ?   ALTER TABLE ONLY public.users DROP CONSTRAINT users_email_key;
       public            postgres    false    216            �           2606    16803    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    216            �           2606    16807    users users_username_key 
   CONSTRAINT     W   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_username_key;
       public            postgres    false    216            �           2606    16852 )   order_items order_items_menu_item_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_menu_item_id_fkey FOREIGN KEY (menu_item_id) REFERENCES public.menu_items(id);
 S   ALTER TABLE ONLY public.order_items DROP CONSTRAINT order_items_menu_item_id_fkey;
       public          postgres    false    4787    218    224            �           2606    16847 %   order_items order_items_order_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;
 O   ALTER TABLE ONLY public.order_items DROP CONSTRAINT order_items_order_id_fkey;
       public          postgres    false    224    4791    222            �           2606    16834    orders orders_table_id_fkey    FK CONSTRAINT     |   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_table_id_fkey FOREIGN KEY (table_id) REFERENCES public.tables(id);
 E   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_table_id_fkey;
       public          postgres    false    220    222    4789            S   �   x�m�;�@����*f���:6&Dcgs��!���D�d�ؘT&8��k���R�h4z�Z���@.X�����9GՓӼU=o�W�h|�H�U��XV`k���R�=9��Q�c����Czrsx 	N7�9J
9���w�h2W�7W���WV�5�nh%u�L)Kv�c�gV{      Y      x������ � �      W      x�3�4�4�30�L����� �T      ]   _   x�u�K
�  еs�.��GC<K�
�ɰ�t���#CF����>�z��W;��-Z�a�`�[�)�� n�H���r�y��%&$GAȏ09 x�H#�      [   Q   x�3�J-�/*IUHIUK�+I,�t-rS�J����%�99�)� e9��
e`�@^�Bnj�����������1W� 8��      U      x�3���L*J�2���P�J��qqq �S	�      Q   �   x�e��
�@���S�pm)B-D�hUqm�LtFf4�7j=�/�d�hw��|pm����|zHvG��=6B���:�f.P	�d�؂R���=��;��`�ӫ�8�@���O��.
@�BD4�>�;��(@��;!�t����"�h���Z��^�O{�\����I��e���6.�aw$sP��7������jP�     