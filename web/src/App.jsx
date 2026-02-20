import { useState } from 'react';
import { 
  Box, 
  Drawer, 
  AppBar, 
  Toolbar, 
  Typography, 
  List, 
  ListItem, 
  ListItemButton, 
  ListItemIcon,
  ListItemText, 
  Button, 
  Table, 
  TableBody, 
  TableCell, 
  TableContainer, 
  TableHead, 
  TableRow, 
  Paper,
  Divider,
  Container,
  IconButton,
  useTheme,
  useMediaQuery,
  Grid,
  alpha
} from '@mui/material';
import { BarChart } from '@mui/x-charts/BarChart';
import CloudUploadIcon from '@mui/icons-material/CloudUpload';
import MenuIcon from '@mui/icons-material/Menu';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import DashboardIcon from '@mui/icons-material/Dashboard';
import EuroIcon from '@mui/icons-material/Euro';
import BoltIcon from '@mui/icons-material/Bolt';
import FavoriteIcon from '@mui/icons-material/Favorite';
import CasinoIcon from '@mui/icons-material/Casino';

const drawerWidth = 240;

const games = [
  { name: 'Thunderball', icon: <BoltIcon /> },
  { name: 'EuroMillions', icon: <EuroIcon /> },
  { name: 'Set For Life', icon: <FavoriteIcon /> },
  { name: 'Lotto', icon: <CasinoIcon /> }
];

// Mock data for initial display
const frequencyData = [
  { value: 12, label: '1' },
  { value: 19, label: '2' },
  { value: 3, label: '3' },
  { value: 5, label: '4' },
  { value: 2, label: '5' },
  { value: 14, label: '6' },
];

const drawHistory = [
  { id: 1, n1: 10, n2: 15, n3: 22, cb: 5 },
  { id: 2, n1: 1, n2: 12, n3: 31, cb: 12 },
  { id: 3, n1: 5, n2: 9, n3: 14, cb: 3 },
];

function App() {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));
  const isMidSize = useMediaQuery(theme.breakpoints.between('sm', 'lg'));
  
  // Start closed on mobile, open on larger screens
  const [open, setOpen] = useState(!isMobile);
  const [selectedGame, setSelectedGame] = useState('Thunderball');

  const handleDrawerToggle = () => {
    setOpen(!open);
  };

  const drawerContent = (
    <Box sx={{ display: 'flex', flexDirection: 'column', height: '100%' }}>
      <Toolbar sx={{ display: 'flex', alignItems: 'center', justifyContent: 'flex-start', px: [2] }}>
        <DashboardIcon sx={{ color: theme.palette.primary.main, mr: 1 }} />
        <Typography variant="h6" sx={{ fontWeight: 700, color: 'text.primary' }}>
          Lottery Stats
        </Typography>
        {!isMobile && (
          <IconButton onClick={handleDrawerToggle} sx={{ ml: 'auto' }}>
            <ChevronLeftIcon />
          </IconButton>
        )}
      </Toolbar>
      <Divider />
      <List component="nav" sx={{ p: 2 }}>
        <Typography variant="overline" sx={{ px: 2, fontWeight: 700, color: 'text.secondary' }}>
          Lottery Games
        </Typography>
        {games.map((game) => (
          <ListItem key={game.name} disablePadding sx={{ mb: 0.5 }}>
            <ListItemButton 
              selected={selectedGame === game.name}
              onClick={() => {
                setSelectedGame(game.name);
                if (isMobile) setOpen(false);
              }}
              sx={{
                borderRadius: 2,
                '&.Mui-selected': {
                  backgroundColor: alpha(theme.palette.primary.main, 0.1),
                  color: theme.palette.primary.main,
                  '& .MuiListItemIcon-root': {
                    color: theme.palette.primary.main,
                  },
                },
              }}
            >
              <ListItemIcon sx={{ minWidth: 40 }}>
                {game.icon}
              </ListItemIcon>
              <ListItemText 
                primary={game.name} 
                primaryTypographyProps={{ fontSize: '0.875rem', fontWeight: 500 }}
              />
            </ListItemButton>
          </ListItem>
        ))}
      </List>
    </Box>
  );

  return (
    <Box sx={{ display: 'flex', bgcolor: '#f9fafb', minHeight: '100vh' }}>
      <AppBar 
        position="fixed" 
        elevation={0}
        sx={{ 
          zIndex: (theme) => theme.zIndex.drawer + 1,
          bgcolor: 'rgba(255, 255, 255, 0.8)',
          backdropFilter: 'blur(8px)',
          borderBottom: '1px solid',
          borderColor: 'divider',
          color: 'text.primary',
          transition: theme.transitions.create(['width', 'margin'], {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.leavingScreen,
          }),
          ...(open && !isMobile && {
            marginLeft: drawerWidth,
            width: `calc(100% - ${drawerWidth}px)`,
            transition: theme.transitions.create(['width', 'margin'], {
              easing: theme.transitions.sharp,
              duration: theme.transitions.duration.enteringScreen,
            }),
          }),
        }}
      >
        <Toolbar sx={{ justifyContent: 'space-between' }}>
          <Box sx={{ display: 'flex', alignItems: 'center' }}>
            <IconButton
              color="inherit"
              aria-label="toggle drawer"
              edge="start"
              onClick={handleDrawerToggle}
              sx={{ mr: 2, ...(open && !isMobile && { display: 'none' }) }}
            >
              <MenuIcon />
            </IconButton>
            {!isMobile && (
              <Button
                variant="contained"
                startIcon={<CloudUploadIcon />}
                sx={{ 
                  borderRadius: 2, 
                  textTransform: 'none',
                  px: 3,
                  boxShadow: theme.shadows[2],
                }}
              >
                Upload Data
              </Button>
            )}
          </Box>
        </Toolbar>
      </AppBar>
      
      <Drawer
        variant={isMobile ? "temporary" : "persistent"}
        open={open}
        onClose={handleDrawerToggle}
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
            borderRight: '1px solid',
            borderColor: 'divider',
            elevation: 0
          },
        }}
      >
        {drawerContent}
      </Drawer>

      <Box 
        component="main" 
        sx={{ 
          flexGrow: 1, 
          p: { xs: 2, sm: 3, md: 4 }, 
          width: '100%',
          transition: theme.transitions.create('margin', {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.leavingScreen,
          }),
          ...(!isMobile && {
            marginLeft: open ? 0 : `-${drawerWidth}px`,
          }),
          ...(!isMobile && open && {
            transition: theme.transitions.create('margin', {
              easing: theme.transitions.easing.easeOut,
              duration: theme.transitions.duration.enteringScreen,
            }),
          }),
          mt: 8
        }}
      >
        <Container 
          maxWidth={isMidSize ? "md" : "lg"} 
          sx={{ 
            px: { xs: 0, sm: 2 },
            transition: 'max-width 0.3s' 
          }}
        >
          <Box sx={{ 
            mb: 4, 
            display: 'flex', 
            flexDirection: { xs: 'column', sm: 'row' }, 
            justifyContent: 'space-between', 
            alignItems: { xs: 'stretch', sm: 'flex-end' }, 
            gap: 2 
          }}>
            <Box>
              <Typography variant="body2" color="text.secondary" sx={{ mb: 0.5 }}>
                Dashboard / Lottery Stats
              </Typography>
              <Typography 
                variant="h4" 
                sx={{ 
                  fontSize: { xs: '1.75rem', sm: '2.125rem', lg: '2.5rem' },
                  fontWeight: 700 
                }}
              >
                {selectedGame}
              </Typography>
            </Box>
          </Box>

          <Grid container spacing={3}>
            {/* Chart Section */}
            <Grid item xs={12}>
              <Paper 
                elevation={0}
                sx={{ 
                  p: { xs: 2, sm: 3 }, 
                  borderRadius: 3, 
                  border: '1px solid',
                  borderColor: 'divider',
                  overflow: 'hidden'
                }}
              >
                <Typography variant="h6" sx={{ mb: 2, fontWeight: 600 }}>
                  Frequency Analysis
                </Typography>
                <Box sx={{ width: '100%', overflowX: 'auto' }}>
                  <Box sx={{ minWidth: 500 }}>
                    <BarChart
                      xAxis={[{ scaleType: 'band', data: frequencyData.map(d => d.label) }]}
                      series={[{ data: frequencyData.map(d => d.value), color: theme.palette.primary.main }]}
                      height={300}
                      margin={{ top: 10, bottom: 30, left: 40, right: 10 }}
                    />
                  </Box>
                </Box>
              </Paper>
            </Grid>

            {/* Table Section */}
            <Grid item xs={12}>
              <Paper 
                elevation={0}
                sx={{ 
                  borderRadius: 3, 
                  border: '1px solid',
                  borderColor: 'divider',
                  overflow: 'hidden'
                }}
              >
                <Box sx={{ p: 3, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <Typography variant="h6" sx={{ fontWeight: 600 }}>
                    Recent Draws
                  </Typography>
                  <Button size="small" sx={{ textTransform: 'none' }}>View All</Button>
                </Box>
                <TableContainer sx={{ overflowX: 'auto' }}>
                  <Table stickyHeader>
                    <TableHead sx={{ bgcolor: alpha(theme.palette.primary.main, 0.02) }}>
                      <TableRow>
                        <TableCell align="center" sx={{ fontWeight: 700 }}>Number 1</TableCell>
                        <TableCell align="center" sx={{ fontWeight: 700 }}>Number 2</TableCell>
                        <TableCell align="center" sx={{ fontWeight: 700 }}>Number 3</TableCell>
                        <TableCell align="center" sx={{ fontWeight: 700 }}>Bonus (CB)</TableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {drawHistory.map((row) => (
                        <TableRow key={row.id} hover>
                          <TableCell align="center">{row.n1}</TableCell>
                          <TableCell align="center">{row.n2}</TableCell>
                          <TableCell align="center">{row.n3}</TableCell>
                          <TableCell align="center">
                            <Box sx={{ 
                              display: 'inline-block', 
                              px: 1.5, 
                              py: 0.5, 
                              borderRadius: 1, 
                              bgcolor: alpha(theme.palette.secondary.main, 0.1),
                              color: theme.palette.secondary.dark,
                              fontWeight: 600
                            }}>
                              {row.cb}
                            </Box>
                          </TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </TableContainer>
              </Paper>
            </Grid>
          </Grid>
        </Container>
      </Box>
    </Box>
  );
}

export default App;
