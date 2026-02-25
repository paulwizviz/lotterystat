import { useState, useEffect, useCallback } from 'react';
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
  alpha,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  CircularProgress,
  Stack
} from '@mui/material';
import { BarChart } from '@mui/x-charts/BarChart';
import FileUploadIcon from '@mui/icons-material/FileUpload';
import MenuIcon from '@mui/icons-material/Menu';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import DashboardIcon from '@mui/icons-material/Dashboard';
import EuroIcon from '@mui/icons-material/Euro';
import BoltIcon from '@mui/icons-material/Bolt';
import FavoriteIcon from '@mui/icons-material/Favorite';
import CasinoIcon from '@mui/icons-material/Casino';

const drawerWidth = 240;

const games = [
  { 
    name: 'Thunderball', 
    id: 'tball', 
    icon: <BoltIcon />, 
    specialLabel: 'Thunderball',
    endpoints: {
      upload: '/tball/csv',
      ballFreq: '/tball/draw/frequency',
      specialFreq: '/tball/tball/frequency'
    }
  },
  { 
    name: 'EuroMillions', 
    id: 'euro', 
    icon: <EuroIcon />, 
    specialLabel: 'Lucky Star',
    endpoints: {
      upload: '/euro/csv',
      ballFreq: '/euro/draw/frequency',
      specialFreq: '/euro/star/frequency'
    }
  },
  { 
    name: 'Set For Life', 
    id: 'sflife', 
    icon: <FavoriteIcon />, 
    specialLabel: 'Life Ball',
    endpoints: {
      upload: '/sflife/csv',
      ballFreq: '/sflife/draw/frequency',
      specialFreq: '/sflife/lball/frequency'
    }
  },
  { 
    name: 'Lotto', 
    id: 'lotto', 
    icon: <CasinoIcon />, 
    specialLabel: 'Bonus Ball',
    endpoints: {
      upload: '/lotto/csv',
      ballFreq: '/lotto/draw/frequency',
      specialFreq: '/lotto/bonus/frequency'
    }
  }
];

function App() {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));
  
  const [open, setOpen] = useState(!isMobile);
  const [selectedGame, setSelectedGame] = useState(games[0]);
  const [openUploadDialog, setOpenUploadDialog] = useState(false);
  const [ballFreqData, setBallFreqData] = useState([]);
  const [specialFreqData, setSpecialFreqData] = useState([]);
  const [loading, setLoading] = useState(false);
  const [uploading, setUploading] = useState(false);
  const [selectedFile, setSelectedFile] = useState(null);

  const fetchFrequencies = useCallback(async () => {
    setLoading(true);
    try {
      const [ballRes, specialRes] = await Promise.all([
        fetch(selectedGame.endpoints.ballFreq),
        fetch(selectedGame.endpoints.specialFreq)
      ]);
      
      if (!ballRes.ok || !specialRes.ok) {
        throw new Error('Failed to fetch data from server');
      }

      const balls = await ballRes.json();
      const specials = await specialRes.json();
      
      console.log(`Fetched ${balls.length} balls and ${specials.length} specials for ${selectedGame.name}`);

      const filteredBalls = balls
        .map(b => ({ value: b.Frequency, label: b.Ball?.toString() }))
        .filter(b => b.value > 0);
        
      const filteredSpecials = specials
        .map(s => ({ 
          value: s.Frequency, 
          label: (s.TBall || s.Star || s.LBall || s.Ball)?.toString() 
        }))
        .filter(s => s.value > 0);

      setBallFreqData(filteredBalls);
      setSpecialFreqData(filteredSpecials);
    } catch (err) {
      console.error('Failed to fetch frequencies:', err);
      setBallFreqData([]);
      setSpecialFreqData([]);
    } finally {
      setLoading(false);
    }
  }, [selectedGame]);

  useEffect(() => {
    fetchFrequencies();
  }, [fetchFrequencies]);

  const handleDrawerToggle = () => {
    setOpen(!open);
  };

  const handleUploadClick = () => {
    setOpenUploadDialog(true);
  };

  const handleUploadClose = () => {
    setOpenUploadDialog(false);
    setSelectedFile(null);
  };

  const handleFileChange = (event) => {
    if (event.target.files && event.target.files[0]) {
      setSelectedFile(event.target.files[0]);
    }
  };

  const handleUploadSubmit = async () => {
    if (!selectedFile) return;

    setUploading(true);
    const formData = new FormData();
    formData.append('file', selectedFile);

    try {
      const response = await fetch(selectedGame.endpoints.upload, {
        method: 'POST',
        body: formData,
      });

      if (response.ok) {
        handleUploadClose();
        fetchFrequencies();
      } else {
        console.error('Upload failed');
      }
    } catch (err) {
      console.error('Error uploading file:', err);
    } finally {
      setUploading(false);
    }
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
              selected={selectedGame.name === game.name}
              onClick={() => {
                setSelectedGame(game);
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
          maxWidth="xl" 
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
              <Typography 
                variant="h4" 
                sx={{ 
                  fontSize: { xs: '1.75rem', sm: '2.125rem', lg: '2.5rem' },
                  fontWeight: 700 
                }}
              >
                {selectedGame.name}
              </Typography>
            </Box>
            <Button
              variant="contained"
              startIcon={<FileUploadIcon />}
              onClick={handleUploadClick}
              sx={{ 
                borderRadius: 2, 
                textTransform: 'none',
                px: 3,
                py: 1,
                boxShadow: theme.shadows[2],
              }}
            >
              Upload Data
            </Button>
          </Box>

          {loading ? (
            <Box sx={{ display: 'flex', justifyContent: 'center', my: 10 }}>
              <CircularProgress />
            </Box>
          ) : (
            <Stack spacing={4} direction="column">
              {/* Main Ball Chart Section */}
              <Paper 
                elevation={0}
                sx={{ 
                  p: { xs: 2, sm: 3 }, 
                  borderRadius: 3, 
                  border: '1px solid',
                  borderColor: 'divider',
                  overflow: 'hidden',
                  width: '100%'
                }}
              >
                <Typography variant="h6" sx={{ mb: 2, fontWeight: 600 }}>
                  Main Ball Frequency Analysis
                </Typography>
                <Box sx={{ width: '100%', height: 400 }}>
                  <BarChart
                    xAxis={[{ scaleType: 'band', data: ballFreqData.map(d => d.label) }]}
                    series={[{ data: ballFreqData.map(d => d.value), color: theme.palette.primary.main }]}
                    height={400}
                    margin={{ top: 10, bottom: 30, left: 40, right: 10 }}
                  />
                </Box>
              </Paper>

              {/* Special Ball Chart Section */}
              <Paper 
                elevation={0}
                sx={{ 
                  p: { xs: 2, sm: 3 }, 
                  borderRadius: 3, 
                  border: '1px solid',
                  borderColor: 'divider',
                  overflow: 'hidden',
                  width: '100%'
                }}
              >
                <Typography variant="h6" sx={{ mb: 2, fontWeight: 600 }}>
                  {selectedGame.specialLabel} Frequency Analysis
                </Typography>
                <Box sx={{ width: '100%', height: 400 }}>
                  <BarChart
                    xAxis={[{ scaleType: 'band', data: specialFreqData.map(d => d.label) }]}
                    series={[{ data: specialFreqData.map(d => d.value), color: theme.palette.secondary.main }]}
                    height={400}
                    margin={{ top: 10, bottom: 30, left: 40, right: 10 }}
                  />
                </Box>
              </Paper>

              {/* Table Section */}
              <Paper 
                elevation={0}
                sx={{ 
                  borderRadius: 3, 
                  border: '1px solid',
                  borderColor: 'divider',
                  overflow: 'hidden',
                  width: '100%'
                }}
              >
                <Box sx={{ p: 3, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <Typography variant="h6" sx={{ fontWeight: 600 }}>
                    Frequency Analysis Table
                  </Typography>
                  <Button size="small" sx={{ textTransform: 'none' }}>View All</Button>
                </Box>
                <TableContainer sx={{ overflowX: 'auto' }}>
                  <Table stickyHeader>
                    <TableHead sx={{ bgcolor: alpha(theme.palette.primary.main, 0.02) }}>
                      <TableRow>
                        <TableCell align="center" sx={{ fontWeight: 700 }}>Ball</TableCell>
                        <TableCell align="center" sx={{ fontWeight: 700 }}>Frequency</TableCell>
                        <TableCell align="center" sx={{ fontWeight: 700 }}>{selectedGame.specialLabel}</TableCell>
                        <TableCell align="center" sx={{ fontWeight: 700 }}>Frequency</TableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {Array.from({ length: Math.max(ballFreqData.length, specialFreqData.length) }).map((_, index) => (
                        <TableRow key={index} hover>
                          <TableCell align="center">{ballFreqData[index]?.label || '-'}</TableCell>
                          <TableCell align="center">{ballFreqData[index]?.value ?? '-'}</TableCell>
                          <TableCell align="center">
                            {specialFreqData[index] ? (
                              <Box sx={{ 
                                display: 'inline-block', 
                                px: 1.5, 
                                py: 0.5, 
                                borderRadius: 1, 
                                bgcolor: alpha(theme.palette.secondary.main, 0.1),
                                color: theme.palette.secondary.dark,
                                fontWeight: 600
                              }}>
                                {specialFreqData[index].label}
                              </Box>
                            ) : '-'}
                          </TableCell>
                          <TableCell align="center">{specialFreqData[index]?.value ?? '-'}</TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </TableContainer>
              </Paper>
            </Stack>
          )}
        </Container>
      </Box>
      <Dialog open={openUploadDialog} onClose={handleUploadClose} maxWidth="sm" fullWidth>
        <DialogTitle>Upload Data for {selectedGame.name}</DialogTitle>
        <DialogContent>
          <DialogContentText sx={{ mb: 2 }}>
            Select a CSV file from your local machine to update the draw history for {selectedGame.name}.
          </DialogContentText>
          <Box 
            component="label"
            sx={{ 
              border: '2px dashed', 
              borderColor: selectedFile ? theme.palette.primary.main : 'divider', 
              borderRadius: 2, 
              p: 4, 
              textAlign: 'center',
              bgcolor: alpha(theme.palette.primary.main, 0.02),
              cursor: 'pointer',
              display: 'block',
              '&:hover': {
                bgcolor: alpha(theme.palette.primary.main, 0.05),
                borderColor: theme.palette.primary.main
              }
            }}
          >
            <input
              type="file"
              accept=".csv"
              hidden
              onChange={handleFileChange}
            />
            <FileUploadIcon sx={{ fontSize: 48, color: selectedFile ? theme.palette.primary.main : 'text.secondary', mb: 1 }} />
            <Typography variant="body1" fontWeight={500}>
              {selectedFile ? selectedFile.name : 'Click to select a CSV file'}
            </Typography>
            <Typography variant="caption" color="text.secondary">
              Support for single CSV file upload.
            </Typography>
          </Box>
        </DialogContent>
        <DialogActions sx={{ px: 3, pb: 3 }}>
          <Button onClick={handleUploadClose} disabled={uploading} sx={{ textTransform: 'none' }}>Cancel</Button>
          <Button 
            onClick={handleUploadSubmit} 
            variant="contained" 
            disabled={!selectedFile || uploading}
            sx={{ textTransform: 'none', px: 3, minWidth: 100 }}
          >
            {uploading ? <CircularProgress size={24} color="inherit" /> : 'Upload'}
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}

export default App;
